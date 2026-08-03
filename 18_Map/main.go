package main

import "fmt"

func main() {
	//map[keyType]ValueType
	age := map[string]int{
		"ghan": 21,
		"sus":  22,
	}
	fmt.Println("Ages: ", age)
	fmt.Println("Age of ghan: ", age["ghan"])
	fmt.Println("Length of age: ", len(age))

	//make to create map
	//make(map[k]V)

	var scores map[string]int //nil map

	scores = make(map[string]int)
	scores["math"] = 90
	scores["science"] = 99
	scores["Python"] = 100

	fmt.Println(scores)
	fmt.Println(scores["math"])
	fmt.Println(scores["science"])
	fmt.Println(scores["Python"])

}
