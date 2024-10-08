package main

import "fmt"

// returns a function, which in turn returns an integer
func intSeq() func() int {
	i := 0
	// the anonymous function closes over the variable i,i.e. the variable i is captured, creating a closure
	// anonymous function - inline
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq() // the closure, now i is bound to nextInt, each time it is called, i gets incremented
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
