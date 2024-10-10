package main

import "fmt"

func multiReturn() (int, int) {
	return 3, 7 // mutiple returns as sepcified above
}

func main() {
	res1, res2 := multiReturn()
	fmt.Println(res1, res2)

}
