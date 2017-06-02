package main

import "fmt"

// Storer is an interface for things that can store
// cargo
type Storer interface {
	Stow(string)
	Inventory() []string
}

type Subaru struct {
	Trunk []string
}

func (s *Subaru) Stow(item string) {
	s.Trunk = append(s.Trunk, item)
}

func (s *Subaru) Inventory() []string {
	return s.Trunk
}

type Boeing struct {
	CargoHold []string
}

func (b *Boeing) Stow(item string) {
	b.CargoHold = append(b.CargoHold, item)
}

func (b *Boeing) Inventory() []string {
	return b.CargoHold
}

func main() {
	s := &Subaru{}
	b := &Boeing{}

	PutCargo(s, "spare tire")
	PutCargo(b, "suitcase")

	GetInventory(s)
	GetInventory(b)

}

func PutCargo(s Storer, item string) {
	s.Stow(item)
}

func GetInventory(s Storer) {
	fmt.Println(s.Inventory())
}
