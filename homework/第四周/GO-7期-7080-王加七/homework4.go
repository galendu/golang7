package main

import (
	"errors"
	"fmt"
)

//1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
func daoshu(f ...float64) (float64, error) {
	sum := 1.
	for _, v := range f {
		sum *= v
	}
	if sum == 0 {
		return -1, errors.New("divide by 0 ")
	} else {
		return 1 / sum, nil
	}
}

//2.上题用递归实现
func div(f ...float64) (float64, error) {
	if len(f) == 0 { //若为空，做判断，确保有首元素
		return 0, errors.New("divide by 0")
	}

	if f[0] == 0 {
		return 0, errors.New("divide by 0")
	}

	if len(f) == 1 { //当剩余部分仅为1，直接返回结果
		return 1 / f[0], nil
	}

	remain := f[1:]
	if res, err := div(remain...); err != nil {
		return 0, err
	} else {
		return 1 / f[0] * res, nil //假设为3个参数:1，2，3，则思路: 1 / ( div(1)*div(2,3) )
	}
}

//3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
type Fish interface {
	Swim()
	Jump()
}
type Reptile interface {
	Crawl()
}

type Frog struct{}

func (Frog) Swim() {
	fmt.Println("I can swim!")
}

func (Frog) Jump() {
	fmt.Println("I can jump!")
}

func (Frog) Crawl() {
	fmt.Println("I can crawl")
}

//4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
func square(num interface{}) interface{} {
	switch v := num.(type) {
	case float32:
		return v * v
	case float64:
		return v * v
	case int:
		return v * v
	case byte:
		return v * v
	default:
		return "输入类型不正确！"
	}

}

func main() {
	//1
	var a, b = 2., 5.
	if res, err := daoshu(a, b); err != nil {
		fmt.
			Println(err.Error())
	} else {
		fmt.Println(res)
	}

	//2
	if res, err := div(1., 2., 10., 2.5); err != nil {
		fmt.
			Println(err.Error())
	} else {
		fmt.Println(res)
	}

	//3
	var frog Frog
	var fi Fish
	var re Reptile

	fi = frog
	re = frog

	fi.Jump()
	fi.Swim()
	re.Crawl()

	//4
	fmt.Println(square(3))
	fmt.Println(square(1.0))
	fmt.Println(square(1.1213))
	fmt.Println(square(byte(255)))
	fmt.Println(square(nil))
}
