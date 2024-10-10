// variadic functions can take any number if args

package main

import "fmt"

// x is the name of the iterable(which stores all the args)
func variadic_function(x ...int) int {
	fmt.Println(x, " ")
	// we can use range function over the parameters
	sum := 0
	for i := range x {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(variadic_function(1, 2, 3, 4))
	fmt.Println(variadic_function(2, 3, 5, 6))
}
