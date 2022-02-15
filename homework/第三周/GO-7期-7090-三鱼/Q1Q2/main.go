package main

import (
	"fmt"
	"math/rand"
)

func work1() {
	const (
		ROW    = 8
		COLUMN = 5
	)

	var A = [ROW][COLUMN]int{}
	var B = [ROW][COLUMN]int{}

	//赋值
	for i := 0; i < ROW; i++ {
		for j := 0; j < COLUMN; j++ {
			A[i][j] = rand.Intn(10)
			B[i][j] = rand.Intn(10)
		}
	}

	//求和
	var C = [ROW][COLUMN]int{}
	for i := 0; i < ROW; i++ {
		for j := 0; j < COLUMN; j++ {
			C[i][j] = A[i][j] + B[i][j]
			fmt.Printf("%3d ", C[i][j])
		}
		fmt.Println()
	}
}

func work2_if(num int) string {
	if num > 0 && num <= 3 {
		return "春"
	} else if num > 3 && num <= 6 {
		return "夏"
	} else if num > 6 && num <= 9 {
		return "秋"
	} else if num > 9 && num <= 12 {
		return "冬"
	} else {
		return "输入错误！"
	}
}

func work2_switch(num int) string {
	switch num {
	case 1, 2, 3:
		return "春"
	case 4, 5, 6:
		return "夏"
	case 7, 8, 9:
		return "秋"
	case 10, 11, 12:
		return "冬"
	default:
		return "输入错误！"

	}
}

func main() {
	//1
	work1()

	fmt.Println("========================")

	//2_if
	fmt.Println(work2_if(1))
	fmt.Println(work2_if(6))
	fmt.Println(work2_if(13))
	//2_switch
	fmt.Println(work2_switch(1))
	fmt.Println(work2_switch(6))
	fmt.Println(work2_switch(13))

	fmt.Println("========================")
}

// 逻辑不错。第一题可以创建一个生成矩阵的方法函数
