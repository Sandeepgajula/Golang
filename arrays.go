package main

import "fmt"

func main() {
	var a [4]int
	a[0] = 1
	a[1] = 10
	a[2] = 11
	a[3] = 12
	fmt.Println(a)
	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)
	c := []int{1, 2, 3}
	fmt.Println(len(c))
	fmt.Println(c)
	d := []string{"hii", "hello"}
	fmt.Println(d)

}
