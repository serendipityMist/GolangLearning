package main

import (
	"fmt"
)

func main() {
	view1 := 1000
	view2 := 2000
	total_views := view1 + view2

	likes := 100
	likes++
	likes++

	avgViews := total_views / 2
	fmt.Println("Average View: ", avgViews)
	fmt.Println("Total view: ", total_views)
	fmt.Println("Total likes: ", likes)

	rating1 := 4.5
	rating2 := 5.1

	avgRating := (rating1 + rating2) / 2
	fmt.Println("Average rating is: ", avgRating)

}
