package main

import "fmt"

func main() {

	scores := make([]int, 0, 5)

	scores = append(scores, 100, 200, 300)
	fmt.Println("After appending 100,200,300", scores, len(scores), cap(scores))

	scores = append(scores, 4000, 500, 533)
	fmt.Println("After appending 4000,500,533", scores, len(scores), cap(scores))

	scores = append(scores, 123, 124, 125, 126, 890, 99)
	fmt.Println("After appending more values. To check how much the capcity grows by: ", scores, len(scores), cap(scores))

	todos := []string{"My", "name", "is", "mist"}
	more := []string{"I", "am", "learning", "Golang"}

	//... spreads
	todos = append(todos, more...)
	fmt.Println(todos)
}
