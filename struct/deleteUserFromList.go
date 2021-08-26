package _struct

import (
	"fmt"
	_ "fmt"
)

type UserData struct {
	name string
	age  int
	id   int
}

func DeleteUserFronListOfStruct() {
	listOfUsers := []UserData{
		UserData{name: "Laksh", age: 21, id: 0},
		UserData{name: "Aslam", age: 13, id: 0},
		UserData{name: "Amir", age: 31, id: 0},
		UserData{name: "Asim", age: 25, id: 0},
		UserData{name: "Fakhar", age: 41, id: 0},
		UserData{name: "Sharjeel", age: 46, id: 0},
		UserData{name: "David", age: 29, id: 0},
		UserData{name: "Jeff", age: 32, id: 0},
		UserData{name: "David", age: 29, id: 0},
		UserData{name: "Jeff", age: 32, id: 0},
	}
	fmt.Println("List before deleting Laksh", listOfUsers)
	result := findAndDeleteByName(listOfUsers, "Laksh")
	fmt.Println(result)
}

func findAndDeleteByName(data []UserData, name string) []UserData {
	index := 0

	for _, element := range data {
		if element.name != name {
			data[index] = element
			index++
		}
	}
	return data[:index]
}
