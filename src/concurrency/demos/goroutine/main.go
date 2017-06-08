package main

import "fmt"
import "time"

func main() {
	go sayHello()
}

func sayHello() {
	fmt.Println("hello")
}
