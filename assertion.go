package main

import (
	"fmt"
)

func do(i interface{}) {
	// interface iをint型に変換
	ii := i.(int) * 2
	fmt.Println(ii)
}

func main() {
	do(10)
}
