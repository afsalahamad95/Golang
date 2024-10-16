package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(1 * time.Second)
	// afterfunc will execute the function after the timer expires
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("Function executed after timeout")
	})
	fmt.Println("Timer fired")
	<-timer.C // This blocks the execution of the program until the timer expires
	fmt.Println("Timer expired")
	time.Sleep(2 * time.Second)

	// returns a channel that will send the current time after specified timeout without creating a timer
	<-time.After(1 * time.Second)
	fmt.Println("1 second passed")

	// to reset the timer, use the Reset method
	timer.Reset(2 * time.Second) // timer will now expire in 2 seconds
	<-timer.C
	timer.Stop()
	fmt.Println("Timer stopped")
}
