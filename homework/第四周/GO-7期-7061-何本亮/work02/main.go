package main

import (
	"errors"
	"fmt"
)

func multi(a ...int) int {
	i := len(a)
	if i == 1 {
		return a[i-1]
	}
	return a[i-1] * multi(a[:i-1]...)
}

func reciprocal(i ...int) (float32, error) {
	result := multi(i...)
	if result == 0 {
		return -1, errors.New("divide by zero")
	}
	return 1 / float32(result), nil
}

func main() {
	if res, err := reciprocal(0); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("倒数为: %f\n", res)
	}

}

// 和第一题类似，想一想是不是少了一种条件判断
