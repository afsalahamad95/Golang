package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // to avoid generating the same random number, we're seeding using the unix epoch which will give precise time in nanoseconds, which will enforce new random numbers every time we run it
	diceNumber := rand.Intn(6) + 1
	fmt.Println("current value:", diceNumber)
	switch diceNumber {
	// cases will break automatically
	// in case you don't want to break, add fallthrough statement in the case
	case 1:
		fmt.Println("move one step")
	case 2:
		fmt.Println("move two step")
		// fallthrough
	case 3:
		fmt.Println("move 3 step")
	case 4:
		fmt.Println("move 4 step")
	case 5:
		fmt.Println("move 5 step")
		// fallthrough -> avoids the break statement here

	case 6:
		fmt.Println("move 6 step")
	default:
		fmt.Println("Invalid value")
	}
}
