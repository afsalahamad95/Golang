package main

import (
	"fmt"
	"sync"
)

func x(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go x("hi", &wg)
	go x("hello", &wg)
	go x("world", &wg)
	wg.Wait()
	fmt.Println("done")
}
