package main

import "fmt"

/*
3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
*/
type Fisher interface {
	Swim()
}

type Reptileer interface {
	crawl(int)
}

type Frog struct {
	Name string
}

func (frog Frog) Swim() {
	fmt.Printf("%s是鱼类，它会游泳\n", frog.Name)
}

func (frog Frog) crawl(i int) {
	fmt.Printf("%s是爬行动物，它能爬行%d米\n", frog.Name, i)
}

func Swim(frog Frog, fisher Fisher) {
	fmt.Printf("%s是鱼类，它会游泳\n", frog.Name)
}

func crawl(i int, frog Frog, fisher Fisher) {
	fmt.Printf("%s是爬行动物，它能爬行%d米\n", frog.Name, i)
}

func main() {
	var f Frog
	f = Frog{Name: "青蛙王子"}
	var (
		fisher    Fisher
		reptileer Reptileer
	)
	fisher = f
	reptileer = f
	Swim(f, f)
	crawl(10, f, f)
	fisher.Swim()
	reptileer.crawl(10)
}

// 后面单独的对应的函数不用实现即可
