package main

import "fmt"

/*
4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
*/

func square(num interface{}) interface{} {
	switch a := num.(type) {
	case float32:
		return a * a
	case float64:
		return a * a
	case int:
		return a * a
	case byte:
		return a * a
	default:
		return "请输入float32、float64、int、byte4种类型的参数"
	}
}

func main() {
	fmt.Println(square(10))
	fmt.Println(square(10.10))
	fmt.Println(square(byte(10)))
	fmt.Println(square("10"))
}

// 完成的不错
