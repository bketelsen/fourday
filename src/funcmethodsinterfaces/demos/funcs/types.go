package main

import "fmt"

var c = 10

type user struct {
	name string
}
type vehicle struct {
	name string
}

func Name(u user) string {
	return u.name
}

type Namer interface {
	Name() string
}

func (u user) Name() string {
	return u.name
}

func (v vehicle) Name() string {
	return v.name
}

func main() {
	u := user{name: "Brian"}
	v := vehicle{name: "Tesla"}
	printName(u)
	printName(v)
}

func printName(n Namer) {
	fmt.Println(n.Name())
}
