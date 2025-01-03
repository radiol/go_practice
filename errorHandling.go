package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("errorHandling.go")
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	fmt.Println(count, "bytes read:", string(data))
}
