package loops

import (
	"fmt"
	_ "fmt"
)

func Removespace() {
	str := "hello how are you?"
	for _, element := range str {
		cmp := string(element) == " "
		if cmp {
			continue
		}
		fmt.Printf("%s", string(element))
	}
}
