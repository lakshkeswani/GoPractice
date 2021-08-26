package _struct

import (
	"fmt"
	_ "fmt"
)

type Animal struct {
	Name   string
	Origin string
}
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func ExtentStruct() {
	bird := Bird{
		Animal: Animal{
			Name:   "Parrot",
			Origin: "pakistan",
		},
		SpeedKPH: 20,
		CanFly:   true,
	}
	fmt.Println(bird)
}
