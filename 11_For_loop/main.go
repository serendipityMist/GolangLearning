package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println("No. ", i)
	}

	n := 10
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("The sum of n number is: ", sum)
}
