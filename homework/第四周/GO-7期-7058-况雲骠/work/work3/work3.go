package work3

import "fmt"

type Fesh interface {
	Swim()
}

type Animal interface {
	Run()
}

type Frog struct {
}

func (*Frog) Swim() {
	fmt.Println("我会游泳")
}
func (*Frog) Run() {
	fmt.Println("我会奔跑")
}
