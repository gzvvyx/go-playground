package main

import (
	"fmt"
	"time"
)

func Hello(str string) string {
	time.Sleep(2 * time.Second)
	fmt.Println(str, "World!")
	return str
}

func main() {
	fmt.Println("Start")
	var idk string
	idk = Hello("Hello")
	fmt.Println(idk, "back")
}
