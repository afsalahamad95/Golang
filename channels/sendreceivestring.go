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
		ch1 <- "hello"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "world"
	}()
	res1 := <-ch1
	fmt.Println(res1)
	res2 := <-ch2
	fmt.Println(res2)
}
