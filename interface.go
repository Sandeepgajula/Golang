package main

import (
	"fmt"
)

type Container struct {
	radius float64
	height float64
}
type Area interface {
	Tarea() float64
	Tvolume() float64
}

func (c Container) Tarea() float64 {
	return 3.14 * c.radius * c.radius
}
func (c Container) Tvolume() float64 {
	return 3.14 * c.radius * c.radius * c.height
}

func main() {
	var shape Area = Container{10.2, 22.4}
	fmt.Println(shape.Tarea())
	fmt.Println(shape.Tvolume())
}
