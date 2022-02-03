package main

import (
	"errors"
	"fmt"
)

//实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error，用递归实现

func multiplication(args ...float64) float64 {
	var (
		count int
	)
	count = len(args)
	/*  // 由另一个函数传递的，在另一个函数中已经做了传参为空的判断，所以这里不需要
	if count == 0 { // 传入参数为空
		return 0.0
	}
	*/
	if count == 1 { // 只剩最后一个数，递归出口
		return args[0]
	}

	return args[count-1] * multiplication(args[:count-1]...) // 递归调用
}

func reciprocal(args ...float64) (reciprocal float64, err error) {
	var (
		count  int
		result float64
	)
	count = len(args)
	if count == 0 { // 传入参数为空
		reciprocal = 0.0
		err = errors.New("Parameter can't be empty!")
		return
	}
	for _, num := range args {
		if num == 0 { // 参数中有任何一个为0的数，直接返回
			err = errors.New("Divisor can't be zero!")
			reciprocal = 0.0
			return
		}
	}
	result = multiplication(args...) // 调用递归计算乘积
	if result == 0 {          // 乘积为0
		err = errors.New("Divisor can't be zero!")
		reciprocal = 0.0
		return
	}
	reciprocal = 1 / result
	return
}

func main() {
	var (
		a      = 3.0
		b      = 12.0
		c      = 14.0
		result float64
		err    error
	)
	if result, err = reciprocal(a, b, c); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
