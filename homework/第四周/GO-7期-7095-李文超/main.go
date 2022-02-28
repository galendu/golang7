package main

import (
	"errors"
	"fmt"
)

/*
1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
2.上题用递归实现
3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
*/

func main() {
	//1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
	list := []float64{1.0, 3.0}
	alex, err := main1(list...)
	if err == nil {
		fmt.Printf("result 的结果是 %.2f\n", alex)
	} else if err != nil {
		//	fmt.Println(err)
	}

	//2.上题用递归实现
	fmt.Println(digui(32.0))

	//3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
	inet := Qingw{}
	fmt.Println(inet.swim())
	fmt.Println(inet.run())

	//4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
	fmt.Println(square(4))
}

//1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
func main1(f ...float64) (float64, error) {
	if len(f) < 1 {
		return 0., errors.New("输入参数个数为0")
	}
	result := 1.
	for _, value := range f {
		if value == float64(0) {
			return 0, errors.New("参数中有0，除数不能为0")
		}
		result = value
	}
	return 1. / result, nil
}

//2.上题用递归实现
func digui(num ...float64) float64 {
	l := len(num)
	if l == 1 {
		return num[0]
	}
	return num[l-1] * digui(num[:l-1]...)

}

//3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
type (
	Fish interface {
		swim() string
	}
	Papa interface {
		run() string
	}
)
type Qingw struct {
	name  string
	speak string
}

func (q Qingw) swim() string {
	q.speak = "呱呱"
	return q.speak
}
func (q Qingw) run() string {
	q.name = "fire"
	return q.name
}

//4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
func square(num interface{}) interface{} {
	switch n := num.(type) {
	case int:
		return n * n
	case float32:
		return n * n
	case float64:
		return n * n
	case byte:
		return n * n
	default:
		return "请输入正确类型"
	}
}
