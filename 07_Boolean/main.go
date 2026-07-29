package main

import (
	"fmt"
)

func main() {
	isLogged := true
	isAdmin := false
	hasSubscription := false

	//AND &&
	canAddItems := isLogged && isAdmin

	//OR ||
	canDeleteItems := isAdmin || (isLogged && hasSubscription)

	//NOT !
	userIsAdmin := !isAdmin

	//printing the above data
	fmt.Println("AND Case: ", canAddItems)
	fmt.Println("OR Case: ", canDeleteItems)
	fmt.Println("NOT Case: ", userIsAdmin)

}
