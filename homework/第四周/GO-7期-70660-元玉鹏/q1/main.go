package main

import (
	"errors"
	"fmt"
)

func inverse(element ...float64) (float64, error) {

	if len(element) == 0 {
		return 0, errors.New(" division by zero")
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
