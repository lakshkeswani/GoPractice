package interfaces

import "fmt"

type animalInterface interface {
	breathe()
	walk()
}

type human interface {
	animalInterface
	speak()
}

type employee struct {
	name string
}

func (e employee) breathe() {
	fmt.Println("Employee breathes")
}

func (e employee) walk() {
	fmt.Println("Employee walk")
}

func (e employee) speak() {
	fmt.Println("Employee speaks")
}

func EmbeddingInterface() {
	var h human
	h = employee{name: "John"}
	h.breathe()
	h.walk()
	h.speak()
}
