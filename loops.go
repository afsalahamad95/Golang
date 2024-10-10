package main

import "fmt"

func main() {
	days := []string{"moday", "tue", "wed", "thu"}
	fmt.Println(days)
	// usual for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// using range
	for index, day := range days {
		fmt.Println(index, ":", day)
	}

	ptr := 0
	for ptr < 4 {
		if ptr == 2 {
			goto dx // the goto statement, what happens is, the control moves to the label, and execution continues from there
		}
		fmt.Println(days[ptr])
		ptr++
	}
	// goto statement
	// create a label here
dx:
	fmt.Println("Jumped to the label")
}
