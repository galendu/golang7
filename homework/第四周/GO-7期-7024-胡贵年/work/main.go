package main

import (
	"errors"
	"fmt"
)

// 1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
// 2.上题用递归实现
// 3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte

//作业1
// 1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
func multip(arr ...float64) (float64, error) {
	var product float64 = 1.
	for _, ele := range arr {
		product *= ele
	}
	if product == 0 {
		return product, errors.New("divide by zero")
	}
	return 1 / product, nil
}

// 2.用递归实现
func multip2(arr ...float64) (float64, error) {

	if len(arr) == 0 {
		return 0, errors.New("arr  len zero")
	}
	first := arr[0]
	if first == 0 {
		return 0, errors.New("divide by zero")
	}
	if len(arr) == 1 {
		return 1 / first, nil
	}
	brr := arr[1:]
	res, err := multip2(brr...)
	if err != nil {
		return 0, err
	} else {
		return 1 / first * res, nil
	}

}

// 3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
type fish interface {
	swim() int
}

type reptile interface {
	climb() int
}

type Frog struct{}

func (frog Frog) swim(distance *int) int {
	return *distance
}

func (frog Frog) reptile(distance *int) int {
	return *distance
}

// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
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
		fmt.Printf("unsurport data type %T\n", num)
		return nil
	}
}

func main() {
	// 作业1
	ok, err := multip(1.4, 3.0, 56, 6, 87, 3, 36, 8, 5, 213, 45, 3)
	fmt.Println(ok, err)

	//作业2
	ok, err = multip2(1.4, 3.0, 56, 6, 87, 3, 36, 8, 5, 213, 45, 3)
	fmt.Println(ok, err)

	//作业3
	frog := Frog{}
	distance := 5
	fmt.Printf("青蛙游了多少 %d 米\n", frog.swim(&distance))
	fmt.Printf("青蛙爬了多少 %d 米\n", frog.reptile(&distance))

	//作业4
	fmt.Printf("平方是%v", square(""))
}
