package main

import (
	"GoPractice/loops"
	_ "GoPractice/loops"
	"fmt"
	_ "fmt"
)

//global variable
var (
	name string = "laksh"
	age  int    = 21
)

// Rules of variable names
// lowercase variables are for within package uppercase will be expose to other packages
// legnth of variable is equal to life of variable
//keep the acronyms
func main() {
	fmt.Println("2D Array output")
	loops.MultidimenssionalArray()
	fmt.Println("remove spaces")
	loops.Removespace()
	fmt.Println("Sum of Given digit")
	loops.SumOfDigits()
	println("Output of Type Conversion")
	typeconversion()
	println("Output of Contants")
	contants()
	//blocked scope
	println("Output of Explore variable")

	var theURL = "https://google.com"
	println("url =", theURL)
	println("before decleration in main ", name, "age = ", age)
	// name local variable age is global
	var name string = "laksh kumar"

	println("after decleration in main ", name, "age = ", age)
	//this is called shadowing go will use the inner most scope
}
