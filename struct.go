package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func main() {
	v := New(3, 4)
	fmt.Println(v)
	fmt.Println(v.x)
	fmt.Println(v.y)
	fmt.Println(v.Area())
	v.Scale(10)
	fmt.Println(v.x)
	fmt.Println(v.y)
	fmt.Println(v.Area())
}
