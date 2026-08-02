package main

import "fmt"

func main() {

	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	sum := 0

	for i, v := range values {
		sum += v
		fmt.Println("i ", i, " v ", v)
	}
	fmt.Println("The sum is: ", sum)
}
