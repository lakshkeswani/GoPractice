package main

import (
	"fmt"
	_ "fmt"
	_ "math"
)

const  a int16 = 27
func contants() {
	// we can not initalize constants with funtions
	// const  myConst float64 = math.Sin(1.57)
	const  a int = 21
	//shadowing
	fmt.Printf("%v,%T",a,a)
}
