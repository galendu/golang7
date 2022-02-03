package main

import (
	"errors"
	"fmt"
)

// 实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error

func reciprocal(args ...float64) (reciprocal float64, err error) {
	var (
		num     float64
		product = 1.0
	)
	if len(args) == 0 {
		err = errors.New("Parameter can't be empty!")
		return 0.0, err
	}
	//fmt.Printf("%T\n", args)  // []float64
	for _, num = range args {
		product *= num
	}
	if product == 0 {
		err = errors.New("Divisor can't be zero!")
		return
	}
	reciprocal = 1 / product
	return reciprocal, err
}

func main() {
	var (
		a      = 2.0
		b      = 15.0
		c      = 10.0
		result float64
		err    error
	)

	if result, err = reciprocal(a, b, c); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
