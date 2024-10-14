package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 2) // second param is a buffered channel, so redundant messages are stored in buffer, if it is not given, the channel is unbuffered
	// ch <- 5 // assign a value
	// fmt.Println(<-ch)
	// above code will not work
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// <-chan denotes the channel is send-only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch) // read channel
		// fmt.Println(<-ch) // read channel
		// listening to a closed channel will result in 0
		// to make sure channel is open before listening, follow:
		value, isChannelOpen := <-ch
		if isChannelOpen {
			fmt.Println(value)
			fmt.Println(isChannelOpen)
		}
		wg.Done()
	}(ch, wg)
	// // <-chan denotes the channel is receive-only
	// marking them like this will make sure you don't close the channel even before sending data
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5 // assign
		ch <- 6
		// make sure to listen when you send values
		close(ch) // close after use
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
