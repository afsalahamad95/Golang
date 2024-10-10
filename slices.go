package main

import (
	"fmt"
	"strings"
)

func main() {
	a := make([]int, 0, 5) // passed in length and capacity of the slice
	fmt.Println(a)
	fmt.Println(a[0:])

	// [...] allows any number of elements in the array, go itself counts the no. of elements
	b := [...]int{1, 2, 3, 4, 7: 10, 8} // elements from index 4 to 7 are initialized to 0 as we've specified 7:10
	fmt.Println(b)
	s := "hello world i am afsal"
	words := strings.Split(s, " ") // splits the string and returns a slice
	fmt.Println(words)
	words = append(words[1:]) // words slice is also modified
	fmt.Println(words)
}
