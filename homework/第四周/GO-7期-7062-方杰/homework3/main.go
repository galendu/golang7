package main

import "fmt"

// 定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口

// 鱼类
type Fish interface {
	Swim() string // 游
}

// 爬行动物
type Crawler interface {
	Action() string // 行动方式
}

// 青蛙
type Frog struct {
	Name string
}

func (frog *Frog) Swim() string {
	return fmt.Sprintf("I'm a frog, My name is %s, and i cam swim.", frog.Name)
}

func (frog *Frog) Action() string {
	return fmt.Sprintf("I'm a frog, My name is %s, and i cam Jump.", frog.Name)
}

func main() {
	var (
		frog *Frog
	)
	frog = &Frog{Name: "tortoise"}
	fmt.Println(frog.Swim())
	fmt.Println(frog.Action())
}
