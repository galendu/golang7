package main

import (
	"errors"
	"fmt"
)

/*
1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
*/

func GetReciprocal(element ...float64) (float64, error) {
	var (
		sum        float64 = 1.0
		reciprocal float64
		err        error
	)

	for _, value := range element {
		if value == 0 {
			return 0, errors.New("除数不能为0")
		} else {
			sum *= value
		}
	}
	reciprocal = 1 / sum

	if len(element) == 0 {
		err = errors.New("输入为空，请输入参数")
		reciprocal = 0
	}

	return reciprocal, err
}

func main() {
	//第一题
	fmt.Println(GetReciprocal(10.0, 0, 424.12414))
	fmt.Println(GetReciprocal(10.0, 100, 424.12414))
	fmt.Println(GetReciprocal())
}

// 完成的不错