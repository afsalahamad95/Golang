package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Write() {
	content := "hello, world"              // move this content to the file
	file, err := os.Create("./myFile.txt") // creates a new file, will return error if not possible to create -> so use comma ok syntax

	if err != nil {
		// we got an error
		panic(err) // throw the error
	}

	length, err := io.WriteString(file, content) // write content to file -> will return length if write success, else returns error

	if err != nil {
		panic(err)
	}

	fmt.Println("write success, file length:", length)
	defer file.Close() // close after use -> use defer to make sure you close only before return
	Read("./myFile.txt")
}

func Read(filePath string) {
	databyte, err := ioutil.ReadFile(filePath) // data is returned in bytes format, hence call it databyte
	if err != nil {
		panic(err)
	}
	fmt.Println("Read text:", databyte)         // bytes form
	fmt.Println("Read text:", string(databyte)) // natural form

}

func main() {
	Write()
}
