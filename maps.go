package main

import (
	"fmt"
)

func main() {
	map1, map2 := make(map[int]int), make(map[string]int) // define 2 maps
	map1[5] = 7
	map1[7] = 5
	map2["hello"] = 1
	map2["world"] = 2
	test2, val1 := map1[5]
	test, val2 := map2["yo"] // test is assigned the value, if value not present it is assigned 0, and val2 is assigned a boolean, true if key present else false
	fmt.Println(test2, val1) // key present
	fmt.Println(test, val2)  // key not present
	// use maps.Equal(map1,map2) to check if two maps are equal

	// to delete from the map
	fmt.Println(map1)
	delete(map1, 5) // give the map and key to delete
	fmt.Println(map1)

	// loop through the maps - using range
	for key, value := range map1 {
		fmt.Println(key, ":", value)
	}
}
