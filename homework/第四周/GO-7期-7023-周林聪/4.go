package main

import (
	"errors"
	"fmt"
)

//1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
func one4(arr ...float64) (a float64, err error) {
	c := 1.
	for i := 0; i < len(arr); i++ {
		c *= arr[i]
		if c == 0 {
			a = 0
			err = errors.New("不能输入0")
		} else {
			a = 1 / c
			err = nil
		}
	}
	if len(arr) == 0 {
		a = 0
		err = errors.New("输入为空")
	}
	return a, err
}

//2.上题用递归实现
func two4(arr ...float64) (float64, error) {
	if len(arr) == 0 {
		return 0, errors.New("输入为空")
	}
	//先进行0值判断
	if arr[0] == 0 {
		return 0, errors.New("不能输入0")
	}
	//再进行倒数计算
	if len(arr) == 1 {
		return 1 / arr[0], nil
	}
	brr := arr[1:]
	c, err := two4(brr...)
	if err != nil {
		return 0, err
	} else {
		return 1 / (arr[0] * c), nil
	}
}

//3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
type fish interface {
	live(string) string
}
type reptiles interface {
	move(string) string
}
type Flog struct {
	name string
	life int
}

func (Flog) live(string) string {
	return "可以生活在水中"
}

func (Flog) move(string) string {
	return "可以通过爬行运动"
}

func three3() {
	flog := Flog{"青蛙", 5}
	var f fish = flog
	fmt.Println(f.live("生活在哪"))
	var r reptiles = flog
	fmt.Println(r.move("生活在哪"))

}

//4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
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
		return "暂时不支持这种类型,只支持float、int、byte"
	}
}

func main() {
	//第一题
	fmt.Println(one4(1, 3, 0, 2))
	fmt.Println(one4(1.2, 2, 2.5))
	fmt.Println(one4())
	fmt.Println(one4(-1.2, 2, 2.5))
	fmt.Println(one4(1, 2, 2))

	//第2题
	// fmt.Println(two4(1, 20, 0))
	// fmt.Println(two4(1.2, 2, 2.5))
	// fmt.Println(two4())
	// fmt.Println(two4(-1.2, 2, 2.5))
	// fmt.Println(two4(1, 2, 2))

	//第3题
	// three3()

	//第4题
	// fmt.Println(square("sss"))
	// fmt.Println(square(3))
	// fmt.Println(square(2.5))
	// fmt.Println(square(nil))
	// fmt.Println(square(-2.5)) //6.25?

}

//go run 4\4.go
// 变量定义的时候要易于明晰
// 第一题建议把len(arr)判断放到for循环上面
// 第二题实现可以进行优化下
// 三四两道题完成的不错
