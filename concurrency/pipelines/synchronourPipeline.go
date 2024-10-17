package main

import "fmt"

// although channels are synchronous, we can still get all data without buffering
// by using range as it will keep reading from channel until it is closed
func createDataChannel(nums []int) <-chan int {
	res := make(chan int)
	go func() {
		for i := range nums {
			res <- i
		}
		close(res)
		// close channel to signal that no more data will be sent
	}()
	return res
}

func createSquaredchannel(ch <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		//  soon as channel receives data, it is read and squared
		for i := range ch {
			res <- i * i
		}
		close(res)
	}()
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	datachannel := createDataChannel(nums)
	squaredChannel := createSquaredchannel(datachannel)
	for i := range squaredChannel {
		fmt.Println(i)
	}
	fmt.Println("done")
}
