package main

import "fmt"

func main() {
	//array in Golang
	//the array in Golang are fixed sized and they do not grow, they are not dynamic
	var marks [5]int
	marks[0] = 10
	marks[1] = 30
	marks[2] = 40
	marks[3] = 50
	marks[4] = 80

	fmt.Println("Marks: ", marks)
	fmt.Println("Lengths of the Marks Array: ", len(marks))

	//array literal -> Quick way to create a fixed sized array
	num := [5]int{20, 40, 60, 80, 100}
	fmt.Println("Num: ", num)
	fmt.Println("Length of Num Array: ", len(num))
}
