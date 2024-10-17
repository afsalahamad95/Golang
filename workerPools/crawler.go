package main

import (
	"fmt"
	"net/http"
	"time"
)

func crawler(id int, works <-chan string, stats chan<- int) {
	for url := range works {
		time.Sleep(time.Second)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(id, "crawled, status 400", url)
			stats <- 400
		} else {
			fmt.Println(id, "crawled, status", resp.StatusCode, url)
			stats <- 200
		}
	}
}
func main() {
	urls := []string{"https://www.google.com", "https://www.facebook.com", "https://www.twitter.com", "https://www.instagram.com", "https://www.linkedin.com"}
	const workers = 3
	num := len(urls)
	works := make(chan string, num)
	stats := make(chan int, num)
	for i := 0; i < workers; i++ {
		go crawler(i, works, stats)
	}
	for _, url := range urls {
		works <- url
	}
	close(works)
	for i := 0; i < num; i++ {
		<-stats
	}
	close(stats)
	fmt.Println("Execution success")
}
