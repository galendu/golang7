package main

import (
	"errors"
	"fmt"
)

func square(num interface{}) interface{} {
	switch v := num.(type) {
	case float64:
		return v * v
	case float32:
		return v * v
	case int:
		return v * v
	case byte:
		return v * v
	default:
		return errors.New("参数类型异常")
	}
}

func main() {
	fmt.Println(square(12321))
}

// 错误提示要清晰些
