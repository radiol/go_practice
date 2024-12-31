package main

import (
	"fmt"
)

func goroutine(s []int, c chan int) {
	defer close(c)
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine(s, c)
	for i := range c {
		fmt.Println(i)
	}
}
