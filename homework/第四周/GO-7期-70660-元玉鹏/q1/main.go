package main

import (
	"errors"
	"fmt"
)

func inverse(element ...float64) (float64, error) {

	if len(element) == 0 {
		return 0, errors.New(" division by zero") // 此处的错误是否和情况对应
	}

	sum := float64(0)

	for _, sub := range element {
		if sum == float64(0) {
			sum = sub
		} else {
			sum *= sub
		}
	}

	return float64(1) / sum, nil
}

func main() {
	fmt.Println(inverse([]float64{3.1, 4.3, 5.5, 6.66666, 7.000}...))
}

// 逻辑可以进一步优化下
