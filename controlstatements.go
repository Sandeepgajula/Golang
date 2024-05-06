package main

import (
	"fmt"
)

func main() {
	var i int = 1
	if i < 10 {
		fmt.Printf("hiii ")
	} else {
		fmt.Printf("Hello")
	}
	fmt.Println(" ")
	j := 2
	switch j {
	case 1, 4, 5:
		fmt.Println("Case 1")
	case 2, 3, 6:
		fmt.Println("Case 2")
	default:
		fmt.Println("default")

	}

}
