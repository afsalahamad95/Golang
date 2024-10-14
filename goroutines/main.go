package main

import (
	"fmt"
	"sync"
)

func main() {

	var score = []int{0}
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	// goroutines
	wg.Add(1)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("one")
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)
	wg.Add(1)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("two")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)
	wg.Add(1)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("three")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)
	wg.Wait()
	fmt.Println(score)
}
