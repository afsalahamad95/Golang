package main

import (
	"fmt"
	"time"
)

// timers -> do something once after a specific time
// tickers -> do something repeatedly at regular intervals

func main() {
	ticker := time.NewTicker(300 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
}
