// 3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口

package main

import "fmt"

type Fish interface {
	swim()
}

type Crawler interface {
	crawl()
}

type frog struct {
	name string
}

func (w frog) swim() {
	fmt.Printf("%s游泳\n", w.name)
}

func (w frog) crawl() {
	fmt.Printf("%s爬行\n", w.name)
}

// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
func square(num interface{}) interface{} {
	switch v := num.(type) {
	case float32:
		return (v * v)
	case float64:
		return (v * v)
	case int:
		return (v * v)
	case byte:
		return (v * v)
	default:
		return ("不符合")
	}
}

func main() {
	var f Fish
	var c Crawler

	w := frog{name: "蝌蚪"}
	f = w
	c = w
	f.swim()
	c.crawl()

	// 4 
	s := square(4)
	s2 := square(4.5)
	s3 := square(66)
	s4 := square(4.0000)
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}