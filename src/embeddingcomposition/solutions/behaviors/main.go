package main

import "fmt"

// Storer is an interface for things that can store
// cargo
type Storer interface {
	Stow(string)
	Inventory() []string
}
type Bicycle struct {
	Trunk
}

type Subaru struct {
	Trunk
}

type Trunk struct {
	Contents []string
}

func (t *Trunk) Stow(item string) {
	t.Contents = append(t.Contents, item)
}

func (t *Trunk) Inventory() []string {
	return t.Contents
}

type Boeing struct {
	Trunk
}

func main() {
	s := &Subaru{}
	b := &Boeing{}
	c := &Bicycle{}

	PutCargo(s, "spare tire")
	PutCargo(b, "suitcase")
	PutCargo(c, "bread")

	GetInventory(s)
	GetInventory(b)
	GetInventory(c)

}

func PutCargo(s Storer, item string) {
	s.Stow(item)
}

func GetInventory(s Storer) {
	fmt.Println(s.Inventory())
}
