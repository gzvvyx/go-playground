package main

import (
	"fmt"
	"time"
)

func hello(str string) int {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello, World!")
	return 5
}

func main() {
	fmt.Println("Start")
	hello("Hello")
}
