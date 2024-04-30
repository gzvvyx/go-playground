package main

import (
	"fmt"
)

func Hello(str string) string {
	fmt.Println(str, "World!")
	return str
}

func main() {
	fmt.Println("Start")
	var idk string
	idk = Hello("Hello")
	fmt.Println(idk, "back")
}
