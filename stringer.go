package main

import (
	"fmt"
)

type Personal struct {
	Name string
	Age  int
}

// String()を実装することでfmt.Println()での出力をカスタマイズできる
func (p Personal) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

func main() {
	mike := Personal{"Mike", 25}
	fmt.Println(mike)
}
