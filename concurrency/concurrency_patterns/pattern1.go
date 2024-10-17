// for - select loop
package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	// for-select loop
	for {
		select {
		// the channel is used to terminate the loop
		case <-done:
			return
		default:
			fmt.Println("working")
		}
	}
}
func main() {
	done := make(chan bool)
	go doWork(done)
	time.Sleep(3 * time.Second)
	// stop after 3 seconds
	close(done)
	// the close will send a signal to the loop and it will terminate

}
