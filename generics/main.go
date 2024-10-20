package main

import "fmt"

// age is a generic type that can be int or float64
func generic[age int | float64](myAge age) {
	fmt.Println(myAge)
}
func main() {
	var age int = 57
	var age2 float64 = 56.7
	generic(age)
	generic(age2)
}
