package main

import "fmt"

func main() {
	defer fmt.Println("hello 1")
	defer fmt.Println("hello 2")
	defer fmt.Println("hello 3")
	// all of these statements are moved to the stack, then will be executed in lifo order
	fmt.Println("world") // executed before others up there

	for i := 1; i <= 10; i++ {
		defer fmt.Println(i) // will be printed in reverse due to the stack
	} // note these are executed before the other defer statements up there, because defer statements are executed just before returning
}
