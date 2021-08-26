package _struct

import (
	"fmt"
	_ "fmt"
)

type Doctor struct {
	id          int
	name        string
	phoneNumber string
	companions  []string
}

func BasicStruct() {
	aDoctor := Doctor{
		id: 3, name: "Laksh", phoneNumber: "03362238876", companions: []string{
			"X", "Y", "Z",
		},
	}
	fmt.Println(aDoctor)
}
