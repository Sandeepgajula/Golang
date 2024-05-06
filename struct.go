package main

import (
	"fmt"
)

type student struct {
	roll int
	name string
}

func main() {
	san := student{
		10,
		"sandeep",
	}
	fmt.Println(san)
}
