package main

import "fmt"

func main() {
	res := make([]int, 5)
	fmt.Println(res)
	fruits := []string{"apple", "orange", "banana"}
	fruits = append(fruits, "appended_fruit") // returns a reassigned slice
	fmt.Println(fruits)
}
