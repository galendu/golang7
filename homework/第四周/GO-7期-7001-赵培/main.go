package main

import (
	"errors"
	"fmt"
)

//1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
func float01(f ...float64) (float64, error) {

	//求取参数的乘积
	product := 1.0
	for _, ele := range f {
		product *= ele
	}
	//对乘积的倒数进行判断
	if product == 0 {
		return product, errors.New("乘积倒数的除数不能为0")
	}
	return 1 / product, errors.New("")
}

//2.上题用递归实现

//定义一个用来计算float64乘积的递归函数
func product01(args ...float64) float64 {
	n := len(args)
	if n == 1 {
		return args[0]
	}
	return args[n-1] * product01(args[:n-1]...) //调用自身函数返回的值
}

func recursion(args ...float64) (float64, error) {
	result := product01(args...)
	if result == 0 {
		return 0, errors.New("Error: the result of multiplication cannot be 0 when the result is divided")
	}
	return 1 / result, nil

}

//3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口

//定义接口：鱼类动物
type fish interface {
	Say1()
}

//定义接口：爬行动物
type paying interface {
	fish
	Say2()
}

//定义结构体
type animal struct {
	name string
}

func (s *animal) Say1() {
	fmt.Printf("我是%s,我可以在水下游泳\n", s.name)
}

func (s *animal) Say2() {
	fmt.Printf("我是%s,我还可以在陆地上爬行\n", s.name)
}

//4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，
//interface{}允许是4种类型：float32、float64、int、byte

func square(num interface{}) interface{} {
	var res interface{}
	switch i := num.(type) {
	case float32:
		fmt.Printf("num的类型是%T\n", num)
		res = float32(i * i)
	case float64:
		fmt.Printf("num的类型是%T\n", num)
		res = float64(i * i)
	case int:
		fmt.Printf("num的类型是%T\n", num)
		res = int(i * i)
	case byte:
		fmt.Printf("num的类型是%T,%v\n", num, num)
		res = byte(i * i)
	default:
		fmt.Println("只允许计算float32、float64、int、byte该4中类型的数据")
	}
	return res
}

func main() {
	// homework1
	// f1, err := float01(10.0, 20.1, 30.2, 0.0)
	f1, err := float01(10.0, 20.1, 30.2, 40.1)
	fmt.Printf("%f %s\n", f1, err)

	//homework2
	fmt.Println("通过使用递归进行求值如下.........")
	f2, err := float01(10.0, 20.1, 30.2, 40.1)
	fmt.Printf("%f %s\n", f2, err)

	//homework3
	var p paying
	p = &animal{"青蛙"}
	p.Say1()
	p.Say2()

	//homework4
	res := square(5.0)
	fmt.Printf("res的类型%T,结果是%v\n", res, res)
	res = square(5)
	fmt.Printf("res的类型%T,结果是%v\n", res, res)
	res = square(byte('h')) //h=8
	fmt.Printf("res的类型%T,结果是%v\n", res, res)
}
