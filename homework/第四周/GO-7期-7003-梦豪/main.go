package main

import (
	"errors"
	"fmt"
)

//作业
// 1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
// 2.上题用递归实现
// 3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte

//1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error

func float_arg(oter ...float64) (sum float64, err error) {
	product := 1.0
	for _, ele := range oter {
		product *= ele
	}
	if product == 0 {
		err = errors.New("除数不能为0")
	}
	return 1 / product, err
}

//递归实现
//2.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
func float_arg2(args ...float64) float64 {
	n := len(args)
	if n == 1 {
		return args[0]
	}
	return args[n-1] * float_arg2(args[:n-1]...)

}
func float_arg3(args ...float64) (float64, error) {
	result := float_arg2(args...)
	if result == 0 {
		return 0, errors.New("test")
	}
	return 1 / result, nil
}

// 3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口

// 鱼类接口
type Fish interface {
	fish()
}

//爬行类接口
type Reptiles interface {
	Reptiles()
}

//结构体青蛙
type Frog struct {
	name string
}

func (f Frog) fish() {
	fmt.Printf("%s:可以水里游\n", f.name)

}
func (f Frog) Reptiles() {
	fmt.Printf("%s:可以地上爬\n", f.name)
}

//第四题
// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
func square(num interface{}) interface{} {
	switch i := num.(type) {
	case float32:
		return (i * i)
	case float64:
		return (i * i)
	case int:
		return (i * i)
	case byte:
		return (i * i)
	default:
		return "暂时不支持改类型，目前支持 float32 float64,int,byte"

	}

}

func main() {
	//第一题
	fmt.Println(float_arg(10.2, 11.1, 30.0))
	// fmt.Println(float_arg2(10.2, 311.1, 30.0))
	//第二题
	fmt.Println(float_arg3(10.2, 11.1, 30.0))
	//第三题
	var c1 = Frog{name: "青蛙"}
	c1.fish()
	c1.Reptiles()
	//第四题
	fmt.Println(square(4.0))
	fmt.Println(square(3))
	fmt.Println(square(3))
	fmt.Println(square(byte('s')))

}
