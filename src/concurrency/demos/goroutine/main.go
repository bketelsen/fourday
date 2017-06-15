package main

import "fmt"

func main() {
	go sayHello()
}

func sayHello() {
	fmt.Println("hello")
}
