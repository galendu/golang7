package main

import (
	"errors"
	"fmt"
)

func reciprocal(oth ...int) (float32, error) {
	a := 1
	for _, ele := range oth {
		a *= ele
	}
	if a == 0 {
		return -1, errors.New("divide by zero")
	}
	return 1 / float32(a), nil
}

func main() {
	if res, err := reciprocal(2, 3, 0, 5); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("倒数为: %f\n", res)
	}

}
