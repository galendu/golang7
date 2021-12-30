package main

import "fmt"

// 单行注释
/* 多行注释 */
// Add 加法操作
func Add(a int, b int) int {
	return a + b
}

// Sub 加法计算
func Sub(a int, b int) int {
	return a - b
}

// Prod 乘法计算
func Prod(a, b int) int {
	return a * b
}

// Div 除法计算
func Dev(a, b float32) float32 {
	return a / b
}

func main() {
	fmt.Printf("10 + 20 = %d\n", Add(10, 20))
	fmt.Printf("30 - 20 = %d\n", Sub(30, 20))
	fmt.Printf("30 * 20 = %d\n", Prod(30, 20))
	fmt.Printf("30 / 20 = %.2f\n", Dev(30, 20))
	/*
	   计算结果：
	   10 + 20 = 30
	   30 - 20 = 10
	   30 * 20 = 600
	   30 / 20 = 1.50
	*/
}
