package main

import (
	"fmt"
	"math"
)

// 实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte

func square(num interface{}) interface{} {
	switch num.(type) {
	case float32:
		return float32(math.Pow(float64(num.(float32)), 2))
	case float64:
		return math.Pow(num.(float64), 2)
	case int:
		return int(math.Pow(float64(num.(int)), 2))
	case byte:
		return byte(math.Pow(float64(num.(byte)), 2))
	default:
		return "Unsupported data type!"
	}
}

func main(){
	var (
		num1 float32 = 3.1
		num2 float64 = 4.1
		num3 int = 6
		num4 byte = 3
		num5 uint8 = 12  // byte ==> uint8
		num6 uint16 = 11
		num7 int8 = 2
		num8 int16 = 8
		num9 int = 0
		num10 float32 = 0.0
		num11 float64 = 0.0
		num12 byte = 0
		num13 int = -2
		num14 float32 = -2.2
		num15 float64 = -3.2
	)
	fmt.Printf("num1: %T, %v\n", square(num1), square(num1))
	fmt.Printf("num2: %T, %v\n", square(num2), square(num2))
	fmt.Printf("num3: %T, %v\n", square(num3), square(num3))
	fmt.Printf("num4: %T, %v\n", square(num4), square(num4))
	fmt.Printf("num5: %T, %v\n", square(num5), square(num5))
	fmt.Printf("num6: %T, %v\n", square(num6), square(num6))
	fmt.Printf("num7: %T, %v\n", square(num7), square(num7))
	fmt.Printf("num8: %T, %v\n", square(num8), square(num8))
	fmt.Printf("num9: %T, %v\n", square(num9), square(num9))
	fmt.Printf("num10: %T, %v\n", square(num10), square(num10))
	fmt.Printf("num11: %T, %v\n", square(num11), square(num11))
	fmt.Printf("num12: %T, %v\n", square(num12), square(num12))
	fmt.Printf("num13: %T, %v\n", square(num13), square(num13))
	fmt.Printf("num14: %T, %v\n", square(num14), square(num14))
	fmt.Printf("num15: %T, %v\n", square(num15), square(num15))
}
