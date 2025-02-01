package main

import (
	"fmt"
)

type Personal struct {
	Name string
	Age  int
}

// String()を実装することでfmt.Println()での出力をカスタマイズできる
// 実装しない場合、{Mike 25}のように出力される
func (p Personal) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	// Errorを返す時は&と*でポインタレシーバーを返すことを推奨
	return &UserNotFound{Username: "Mike"}
}

// カスタムエラー型を定義する
type UserNotFound struct {
	Username string
}

// Error()を実装することでエラーメッセージをカスタマイズできる
// Errorを実装する際は、ポインタ型で実装する
func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func main() {
	mike := Personal{"Mike", 25}
	fmt.Println(mike)
	err := myFunc()
	if err != nil {
		fmt.Println(err)
	}
}
