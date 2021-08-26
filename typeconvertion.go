package main

import (
	"fmt"
	_ "fmt"
	"strconv"
	_ "strconv"
)

func typeconversion() {
var i int =42
	fmt.Printf("value of j = %V,Type = %T \n",i,i)
    var j float32
	fmt.Printf("value of j = %V,Type = %T \n",j,j)
	j=float32(i)
	fmt.Printf("value of j = %V Type = %T \n",j,j)
  	 var str string ;
	str =strconv.Itoa(i)
	println("int converted into string  ",str)

}