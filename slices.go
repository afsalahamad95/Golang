package main

import "fmt"

func main() {
	a := make([]int, 0, 5)
	fmt.Println(a)
	fmt.Println(a[0:])

	// [...] allows any number of elements in the array, go itself counts the no. of elements
	b := [...]int{1, 2, 3, 4, 7: 10, 8} // elements from index 4 to 7 are initialized to 0 as we've specified 7:10
	fmt.Println(b)
}
