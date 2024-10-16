package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "message 1"
	}() // iife

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "message 1"
	}() // iife

	for i := 0; i < 2; i++ {
		// select waits for all the channels to execute successfully, if no response is obtained, it will switch to default statement or block till it receives one
		select {
		case msg1 := <-ch1:
			fmt.Println("message 1 received", msg1)
		case msg2 := <-ch2:
			fmt.Println("message 2 received", msg2)
		}
	}

}
