package interfaces

import "fmt"

type animalis interface {
	breathe()
	walk()
}

type cat string

func (c cat) breathe() {
	fmt.Println("Cat breathes")
}

func (c cat) walk() {
	fmt.Println("Cat walk")
}

func main() {
	var a animalis

	a = cat("smokey")
	a.breathe()
	a.walk()
}
