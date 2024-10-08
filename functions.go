package main

import "fmt"

// int before the parantheses, indicates the return type
func x(y int) int {
	return y
}

// specify types of similar variables in once
func twoParameters(x, y int) int {
	return (x + y)
}

func threeParameters(x, y, z int) int {
	return (x + y + z)
}

func main() {
	fmt.Println(x(5)) // single parameter function
	fmt.Println(twoParameters(1, 2))
	fmt.Println(threeParameters(1, 2, 3))
}
