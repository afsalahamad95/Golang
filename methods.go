// to create a method on a struct

package main

import "fmt"

func main() {
	user := User{"afsal", 20}
	fmt.Println(user.getAge())
}

type User struct {
	name string
	age  int
}

// an object user of type User will have this method
func (user User) getAge() int {
	return user.age
}
