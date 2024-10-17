package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker", id)
}
func main() {
	// create a waitgroup
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {

		go worker(i, &wg)
		wg.Add(1)
	}
	wg.Wait()

}
