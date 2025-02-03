package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v map[string]int
	// 構造体のフィールドにsync.Mutexを埋め込むことで、構造体のメソッドでロックを使えるようになる
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	// 値を変更する箇所をロックする
	c.mux.Lock()
	// 変更後にロックを解除する
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	// 値を取得する箇所も同様にロックする
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := Counter{v: make(map[string]int)}
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c.Inc("key")
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(c.v, c.Value("key"))
}
