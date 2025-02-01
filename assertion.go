package main

import (
	"fmt"
)

func doInt(i interface{}) {
	// interface iをint型に変換
	ii := i.(int) * 2
	fmt.Println(ii)
}

func doString(i interface{}) {
	// interface iをstring型に変換
	ii := i.(string) + "!"
	fmt.Println(ii)
}

func do(i interface{}) {
	// interface iを型によって処理を変える
	// switch文で型を判定. typeアサーション
	// i.(type)は必ずswitch文の中で使う
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main() {
	doInt(10)
	doString("Mike")
	do(10)
	do("Mike")
	do(true)
}
