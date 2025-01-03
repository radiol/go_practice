package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("defer.go")
	defer file.Close()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println("File data: ", string(data))
}
