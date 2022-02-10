package main

import (
	"fmt"
	"math/rand"
)

// 循环生成[8][5]int 嵌套数组
func matrix(transverse, vertical int) [8][5]int {
	C := [8][5]int{}
	for i := 0; i < transverse; i++ {
		for j := 0; j < vertical; j++ {
			r1 := rand.Intn(100)
			C[i][j] = r1
		}
	}
	return C
}

// 计算两个[8][5]int数组相同位置的和
func sumMatrix(transverse, vertical int, a1, a2 [8][5]int) [8][5]int {
	sum := [8][5]int{}
	for i := 0; i < transverse; i++ {
		for j := 0; j < vertical; j++ {
			sum[i][j] = a1[i][j] + a2[i][j]
		}
	}
	return sum
}

func main() {
	// 两个[8][5]相加的和
	a1 := matrix(8, 5)
	a2 := matrix(8, 5)
	for _, array1 := range a1 {
		fmt.Printf("%5d\n", array1)
	}
	fmt.Println("----------------------------")
	for _, array2 := range a2 {
		fmt.Printf("%5d\n", array2)
	}
	fmt.Println("----------------------------")
	a3 := sumMatrix(8, 5, a1, a2)
	for _, array3 := range a3 {
		fmt.Printf("%5d\n", array3)
	}

}

// 完成的不错
