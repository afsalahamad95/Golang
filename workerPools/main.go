package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, res chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		res <- j
	}
}
func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	res := make(chan int, numJobs)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, res)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	// channels are blocking by default, so receive data from res channel
	for i := 1; i <= numJobs; i++ {
		<-res
	}
}
