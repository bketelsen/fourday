package main

import "fmt"

type thing struct {
	name string
}

func main() {
	t := thing{name: "Brian"}
	fmt.Println(&t)
	fmt.Println(t)
	fmt.Println((&t).name, t.name)
}
