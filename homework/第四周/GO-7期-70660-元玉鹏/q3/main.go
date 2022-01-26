package main

import "fmt"

type Fish interface {
	Swimming()
}

type Reptile interface {
	Breath()
}

type Frog struct {
	Name   string
	Length float32
}

func (f Frog) Swimming() {
	fmt.Printf("%s can Swimming\n", f.Name)
}

func (f Frog) Breath() {
	fmt.Printf("%s can Breath\n", f.Name)

}

func main() {
	var f Frog = Frog{
		Name:   "zhangsan",
		Length: 100.0,
	}

	var fish = f
	var reptile = f
	fish.Swimming()
	reptile.Breath()

}
