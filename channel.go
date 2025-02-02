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

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v * 2
	}
	c <- sum
}

func goroutine3(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		// ループ毎に値を送信
		c <- sum
	}
	// 値の送信が終わったことを示すためにチャネルをクローズ
	close(c)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	// int型のチャネルを作成
	c := make(chan int)
	// goroutineにchannelを渡す
	go goroutine1(s, c)
	go goroutine1(s, c)
	go goroutine2(s, c)
	// x := <-cはchannelから値が取り出されるまでブロックされる
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
	z := <-c
	fmt.Println(z)

	// バッファ付きチャネル, バッファ2のチャネルを作成
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch)) // 1が出力される
	ch <- 200
	fmt.Println(len(ch)) // 2が出力される
	// ch <- 300            // バッファ2を超えるのでエラーになる
	// チャンネルから値を取り出す
	a := <-ch
	fmt.Println(a)
	fmt.Println(len(ch)) // 1が出力される
	// チャンネルから取り出したので入力できる
	ch <- 300
	fmt.Println(len(ch)) // 2が出力される
	fmt.Println("Goroutine3")
	go goroutine3(s, c)
	// rangeでチャネルから値を取り出す
	for i := range c {
		fmt.Println(i)
	}
}
