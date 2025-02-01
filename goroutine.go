package main

import (
	"sync"
	"time"
)

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		println(s, i)
	}
}

func goroutine(s string, wg *sync.WaitGroup) {
	// 処理が終わったらwg.Done()を呼ぶ
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		println(s, i)
	}
}

func main() {
	// WaitGroupを使ってgoroutineの終了を待つ
	wg := sync.WaitGroup{}
	// goroutineを起動する前にAdd(1)を呼ぶ
	wg.Add(1)
	// goroutineに渡す引数はwgのポインタ
	go goroutine("world", &wg)
	normal("hello")
	// goroutineの終了を待つ
	wg.Wait()
}
