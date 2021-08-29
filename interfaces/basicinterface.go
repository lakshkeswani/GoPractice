package interfaces

import (
	"fmt"
	_ "fmt"
	"math"
	_ "math"
)

type Shape interface {
	area() float64
}
type rectangle struct {
	height float64
	width  float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r rectangle) area() float64 {
	return r.height * r.width
}
func BasicInterface() {
	c1 := circle{radius: 4.5}
	r1 := rectangle{height: 5, width: 7}
	Shapes := []Shape{c1, r1}
	fmt.Println("area of Circle = ", Shapes[0].area())
	fmt.Println("area of Rectangle = ", Shapes[1].area())

}
