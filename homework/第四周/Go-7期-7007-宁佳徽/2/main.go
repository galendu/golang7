package main

import (
	"errors"
	"fmt"
)

/*
第二题，用递归实现第一题
*/
func Recursion(element ...float64) float64 {

	count := len(element)

	if count == 1 {
		return element[0]
	}

	//element[count - 1] 等于函数这一循环的值，要获得总值
	//则需要乘以之前函数的值element[:count -1 ], 切片，前包后不包，假如count=3，则element[:count -1 ]==element[0:1]
	return element[count-1] * Recursion(element[:count-1]...)
}

func GetReciprocal2(element ...float64) (float64, error) {
	var (
		reciprocal float64
		err        error
	)

	if len(element) == 0 {
		err = errors.New("输入为空，请输入参数")
		reciprocal = 0
		return reciprocal, err
	}

	for _, value := range element {
		if value == 0 {
			err = errors.New("除数不能为0")
			reciprocal = 0.0
			return reciprocal, err
		}
	}

	result := Recursion(element...)
	if result == 0 {
		err = errors.New("除数不能为0")
		reciprocal = 0.0
		return reciprocal, err
	}
	return 1 / result, err
}

func main() {
	fmt.Println(GetReciprocal2(10.0, 100, 424.12414))
	fmt.Println(GetReciprocal2(10.0, 0, 424.12414))
	fmt.Println(GetReciprocal2())
}
