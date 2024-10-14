package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	// ch <- 5 // assign a value
	// fmt.Println(<-ch)
	// above code will not work
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch) // read channel
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 5 // assign
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
