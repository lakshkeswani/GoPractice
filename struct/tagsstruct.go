package _struct

import (
	"encoding/json"
	_ "fmt"
	"log"
	_ "reflect"
)

type Employee struct {
	EmployeeId int    `json:"employeeid"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
}
type Employee2 struct {
	EmployeeId int
	Name       string
	Age        int
}

func tagstruct() {
	emplyoee1 := Employee{
		EmployeeId: 0,
		Name:       "Laksh",
		Age:        21,
	}
	emplyoee2 := Employee{
		EmployeeId: 1,
		Name:       "Laksh",
		Age:        21,
	}
	j, error := json.Marshal(emplyoee1)
	if error != nil {
		log.Fatal("error by employee ", error.Error())
	}
	println("employee with Tags", j)

	e2, err := json.Marshal(emplyoee2)
	if error != nil {
		log.Fatal("error by employee ", err.Error())
	}

	println("employee without tags", e2)
}
func IsEqualStruct() {
	emplyoee1 := Employee{
		EmployeeId: 0,
		Name:       "Laksh",
		Age:        21,
	}
	emplyoee2 := Employee{
		EmployeeId: 0,
		Name:       "Laksh",
		Age:        21,
	}
	if emplyoee2 == emplyoee1 {
		println("both Employee structs  are equal")
	}
}
