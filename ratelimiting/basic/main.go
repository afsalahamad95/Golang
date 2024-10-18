package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func work(id int, limiter <-chan time.Time) {
	<-limiter
	defer wg.Done()
	fmt.Println("executing", id)
}

func main() {
	limiter := time.Tick(500 * time.Millisecond)
	// limiter to send events every 500 milliseconds
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go work(i, limiter)
	}
	wg.Wait()
}
