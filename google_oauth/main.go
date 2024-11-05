package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleClientID     string `json:"google_client_id"`
	GoogleClientSecret string `json:"google_client_secret"`
}

var (
	config      *Config
	oauthConfig *oauth2.Config
	oauthState  = "random-state-string"
)

func init() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("error occurred when parsing the json file")
	}
	config = &Config{}
	err = json.Unmarshal(file, config)
	if err != nil {
		fmt.Println("an error occurred when reading the parsed json")
	}

	oauthConfig = &oauth2.Config{
		ClientID:     config.GoogleClientID,
		ClientSecret: config.GoogleClientSecret,
		RedirectURL:  "http://localhost:4000/callback",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
		<body>
			<a href = "/login">Login with google</a>
		</body>
	</html>
	`
	fmt.Fprint(w, html)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != oauthState {
		fmt.Println("invalid oauth state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := oauthConfig.Exchange(r.Context(), r.FormValue("code"))
	if err != nil {
		fmt.Println("failed to exchange token")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	client := oauthConfig.Client(r.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		fmt.Println("failed to get userinfo")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		fmt.Println("failed to decode user")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println(userInfo)
	json.NewEncoder(w).Encode("login success")

	user_email := userInfo["email"]

	type Result struct {
		isAdmin bool
	}

	isAdminCheckUrl := fmt.Sprintf("https://admin.googleapis.com/admin/directory/v1/users/%s", user_email)
	resp, err = client.Get(isAdminCheckUrl)
	if err != nil {
		fmt.Println("error occurred when checking details of user")
		return
	}
	defer resp.Body.Close()
	var res Result
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Println("error during the parse of the response body")
		return
	}

	if res.isAdmin {
		fmt.Println("the user is an admin")
	} else {
		fmt.Println("nah he's not an admin bruh!")
	}

}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	fmt.Println("server up and running at port 4000")
	http.ListenAndServe(":4000", nil)

}
