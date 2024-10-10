package main

import "fmt"

func main() {
	var username string = "hello"
	fmt.Printf("type of the variable is %T", username)  // will print the datatype of username
	fmt.Println("type of the variable is %T", username) // will print %T as is and hello as well
}
