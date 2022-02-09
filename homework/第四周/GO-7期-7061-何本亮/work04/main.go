package main

import "fmt"

func square(num interface{}) interface{} {
	switch v := num.(type) {
	case int:
		return v * v
	case float32:
		return v * v
	case float64:
		return v * v
	case byte:
		return v * v
	default:
		return "不支持的类型"
	}
}

func main() {
	num := 2.1
	s := square(num)
	fmt.Printf("%v的平方为: %v\n", num, s)
}
