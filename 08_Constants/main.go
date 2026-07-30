package main

import (
	"fmt"
)

func main() {

	//const is used for those values that do not change, they are constant

	//untyped constant - no need to specify the data type here
	const appName = "Go Basics"

	//typed constant - need to specify the data type
	const maxUpload int = 25
	const discountedPrice float64 = 10.3

	fmt.Println("App name is: ", appName)
	fmt.Println("Max upload is: ", maxUpload)
	fmt.Println("Discount Price is: ", discountedPrice)

}
