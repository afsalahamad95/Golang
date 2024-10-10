package main

import "fmt"

func main() {
	v := Node{"afsal", 19, "afsal@gmail.com"} // create the struct
	fmt.Println(v)
	fmt.Printf("A detailed print statement : %+v", v)
	// +%v will print values along with the field name
	fmt.Println(v.age) // print specific fields

}

// define a struct

// syntax is type <name_of_struct> struct{}
type Node struct {
	// define the members
	name  string
	age   int
	email string
}
