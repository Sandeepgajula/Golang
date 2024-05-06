package main

import (
	"fmt"
)

func main() {
	i := 5
	for j := 0; j < i; j++ {
		fmt.Println(j)
	}
	s := "hello"
	for k, v := range s {
		fmt.Println("Index", k, "Value", v)
	}

}
