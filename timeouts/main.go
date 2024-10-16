package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1) // bufferred channel
	// goroutine
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "request"
	}()

	// timeout using select
	select {
	case msg := <-ch1:
		fmt.Println("200 OK", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Operation timed out")
	}

}
