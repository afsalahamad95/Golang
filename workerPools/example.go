package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, res chan<- int) {
	for i := range jobs {
		fmt.Println(id, "executing", i)
		time.Sleep(time.Second)
		res <- i
	}
}
func main() {
	const numjobs = 5
	jobs := make(chan int, numjobs)
	res := make(chan int, numjobs)
	// fire the workers
	for i := 0; i < 3; i++ {
		go worker(i, jobs, res)
	}
	// send jobs to channel
	for j := 0; j < numjobs; j++ {
		jobs <- j
	}
	// get jobs
	for j := 0; j < numjobs; j++ {
		<-res
	}
}
