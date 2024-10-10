// to create a method on a struct

package main

import "fmt"

func main() {
	user := User{"afsal", 20}
	fmt.Println(user.getAge())
	user.alterName("ahamad")
	fmt.Println(user, "not changed")
	var userptr *User = &user
	userptr.rootAlterName("afsal ahamad") // altered because of the pointer
	fmt.Println(user)                     // altered user
}

type User struct {
	name string
	age  int
}

// an object user of type User will have this method
func (user User) getAge() int {
	return user.age
}

func (user User) alterName(newName string) {
	user.name = newName // didn't change my original age -> because only the copy of object is passed, hence we'll use pointer to sovle this problem
	fmt.Println(user.name, "the new name")
}

func (user *User) rootAlterName(TheName string) {
	user.name = TheName
}
