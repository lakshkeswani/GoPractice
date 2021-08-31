package exceptionhandling

import (
	"fmt"
	_ "fmt"
)

func PanicExp() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
}
