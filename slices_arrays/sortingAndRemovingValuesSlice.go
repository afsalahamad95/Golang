package main

import (
	"fmt"
	"sort"
)

func main() {
	s := make([]int, 5)
	s[0] = 567
	s[1] = 847
	s[2] = 123
	s[3] = 45
	s[4] = 90
	// sort the array
	isSorted := sort.IntsAreSorted(s) // check if sorted
	fmt.Println(isSorted)
	sort.Ints(s) // sorted
	fmt.Println(s)

	// remove a value
	// this can b e done using the append
	// let's delete 123 at index 2 after the sort
	s = append(s[:2], s[3:]...)
	fmt.Println(s)
}
