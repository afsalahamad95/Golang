package main

import "fmt"

func main() {
	// var ptr *int; // an integer pointer
	// var x int = 5
	// var ptr *int = &x
	// fmt.Println(ptr, " is the value of the integer pointer pointing to variable x") // some address
	// var ptrn *int
	// fmt.Println(ptrn, " is the value of a default pointer") // nil

	myNumber := 23
	ptr := &myNumber // point to variable
	*ptr = 35
	*ptr = *ptr * 2
	fmt.Println(myNumber) // now the value is modified

}
