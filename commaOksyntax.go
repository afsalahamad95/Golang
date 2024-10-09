package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("hello world")

	reader := bufio.NewReader(os.Stdin) // we create a reader that reads input from stdin(the keyboard)
	fmt.Println("Enter the input")
	// comma ok syntax - input takes the "ok" value, error is stored in the other variable
	input, err := reader.ReadString('\n') // read the input - \n is the terminator, reading will end if this is given as input
	fmt.Println(input, err)
}
