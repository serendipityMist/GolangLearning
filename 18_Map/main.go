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
}
