package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ch := make(chan int, n)
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
	// always close channel before proceeding to avoid deadlocks
	for i := range ch {
		fmt.Println(i)
	}
}
