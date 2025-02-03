package main

import (
	"fmt"
	"time"
)

func goroutine1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(3 * time.Second)
	}
}

func goroutine2(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goroutine1(ch1)
	go goroutine2(ch2)

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			fmt.Println("waiting...")
			time.Sleep(500 * time.Millisecond)
		}
	}

}
