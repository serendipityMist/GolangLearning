package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Mist"
	lastName := "Shrestha"
	fullName := firstName + " " + lastName

	fmt.Println("Full name is: ", fullName)
	fmt.Println(strings.ToUpper(fullName))
}
