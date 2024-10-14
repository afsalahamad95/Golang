package main

import (
	"fmt"
	"net/http"
	"sync"
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

// func main() {
// 	// use go keyword to create threads
// 	// greeter("hello")
// 	// go greeter("world")
// 	// here second one is not printed because we created a thread, but never waited for it to execute
// 	// to make it work, add time.sleep() in the function, which is one of the ways to execute it
// 	go greeter("hello")
// 	greeter("world")
// }

// func greeter(name string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(2 * time.Millisecond)
// 		fmt.Println(name)
// 	}
// }

// create a waitgroup - it is a modified version of time.sleep() , this will wait till all threads in waitgroup are executed -> and these are usually pointers
var wg sync.WaitGroup

var mutex sync.Mutex // mutex variable -> usually a pointer

var signals = []string{"test"}

func main() {
	websites := []string{"https://go.dev", "https://google.com", "https://youtube.com", "https://github.com"}
	for _, site := range websites {
		go getStatusCode(site) // launch threads
		wg.Add(1)
	}
	wg.Wait() // used after adding all threads to ensure main func waits for threads
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {

	defer wg.Done() // report that thread is executed
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("There was a problem")
	} else {
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
	}
	fmt.Printf("%d Status code\n", res.StatusCode)
}
