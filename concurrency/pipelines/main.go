package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sliceToChannel(nums []int) chan int {
	defer wg.Done()
	ch := make(chan int, len(nums))
	for i := range nums {
		ch <- i
	}
	return ch
}
func square(ch chan int, cap int) chan int {
	defer wg.Done()
	res := make(chan int, cap)
	for i := 0; i < cap; i++ {
		num := <-ch
		res <- num * num
	}
	return res
}
func main() {
	// pipeline is a sequence of data processing stages
	// each stage is a function that takes a channel as an input, processes the data, and sends it to an output channel
	// the output channel of a stage is the input channel to the next stage
	nums := []int{1, 2, 3, 4, 5}
	// stage 1
	wg.Add(1)
	datachannel := sliceToChannel(nums)
	time.Sleep(1 * time.Second)
	// stage 2
	wg.Add(1)
	squaredChannel := square(datachannel, len(nums))
	time.Sleep(1 * time.Second)
	close(datachannel)
	close(squaredChannel)
	for i := range squaredChannel {
		fmt.Println(i)
	}
	wg.Wait()
}
