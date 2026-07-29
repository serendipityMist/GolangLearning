package main

import (
	"fmt"
)

func main() {
	var city = "Kathmandu" //inferred to string

	//short way to declare a variable
	//:=
	subscribers := 5000
	subscribers += 1000

	likes, comments := 100, 30

	fmt.Println("Subscribers: ", subscribers)
	fmt.Println("likes: ", likes)
	fmt.Println("comments	: ", comments)
	fmt.Println("city	: ", city)
}
