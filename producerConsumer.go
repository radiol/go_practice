package main

import (
	"fmt"
	"sync"
	"time"
)

// Producer function
// iを受け取り、chにi*2を送信する
func producer(ch chan int, i int) {
	fmt.Println("Producer", i)
	ch <- i * 2
}

// Consumer function
// chからiを受信し、i*1000を出力する
func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			// mainでwg.Wait()を呼び出すため、wg.Done()を呼び出す
			// mainで10回呼び出すため、10回呼び出す
			defer wg.Done()
			fmt.Println("Consumer", i*1000)
		}()
	}
	fmt.Println("# Consumer Finished")
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}
	// consumerにwgを渡す
	go consumer(ch, &wg)
	// consumerで10回wg.Done()が呼び出されるまで待つ
	wg.Wait()
	// chをクローズすることで、consumerが終了する
	// これをしないと、consumerが終了しない(Consumer Finishedが出力されない)
	close(ch)
	time.Sleep(1 * time.Second)
	fmt.Println("## Main Finished")
}
