package main

import "fmt"

func main() {
	job := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			val, open := <-job
			if open {
				fmt.Println(val, "is received")
			} else {
				done <- true // transfer complete
				fmt.Println("all jobs received")
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		job <- i // send the job
	}
	// closing the channel
	close(job)
	// unblock the main goroutine
	<-done
	_, status := <-job
	fmt.Println("channel status:", status)
}
