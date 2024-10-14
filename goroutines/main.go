package main

import (
	"fmt"
	"time"
)

// normal method -> functions are executed line by line -> no parallelism
// func main() {
// 	greeter("hello")
// 	greeter("world")
// }

// func greeter(name string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(name)
// 	}
// }

func main() {
	// use go keyword to create threads
	// greeter("hello")
	// go greeter("world")
	// here second one is not printed because we created a thread, but never waited for it to execute
	// to make it work, add time.sleep() in the function, which is one of the ways to execute it
	go greeter("hello")
	greeter("world")
}

func greeter(name string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(name)
	}
}
