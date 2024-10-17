package main

import "fmt"

type Employee interface {
	Get() string
}

type Engineer struct {
	name string
}

type Manager struct {
	name string
}

// note that although printdetails takes an employee type, we can pass in a manager type because it implements the employee interface
func (e *Engineer) Get() string {
	return e.name
}
func (m *Manager) Get() string {
	return m.name
}
func PrintDetails(e Employee) {
	fmt.Println(e.Get())
}
func main() {
	engineer := &Engineer{name: "Afsal"}
	PrintDetails(engineer)
	manager := &Manager{name: "John"}
	PrintDetails(manager)
}
