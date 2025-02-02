package main

import (
	"fmt"
)

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	// int型のチャネルを作成
	c := make(chan int)
	// goroutineにchannelを渡す
	go goroutine1(s, c)
	// x := <-cはchannelから値が取り出されるまでブロックされる
	x := <-c
	fmt.Println(x)
}
