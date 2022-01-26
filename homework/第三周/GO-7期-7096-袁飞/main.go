package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
	2.给定月份，判断属于哪个季节。分别用if和switch实现
	3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
*/

// 1. 随机初始化两个8*5的矩阵，求两个矩阵的和(逐元素相加)
func matrix_create() [8][5]int {
	var arr = [8][5]int{}
	for row, array := range arr { // 先取出某一行
		for col, _ := range array { // 再遍历这一行
			arr[row][col] = rand.Intn(100)
		}
	}
	return arr
}

func SumMatrix(x [8][5]int, y [8][5]int) [8][5]int {
	arr := [8][5]int{}
	for row, array := range arr {
		for col, _ := range array {
			arr[row][col] = x[row][col] + y[row][col]
		}
	}

	return arr
}

// 2.给定月份，判断属于哪个季节。分别用if和switch实现
func SeasonIf(x int) {
	if x >= 1 && x <= 3 {
		fmt.Println("春季")
	} else if x >= 4 && x <= 6 {
		fmt.Println("夏季")
	} else if x >= 7 && x <= 9 {
		fmt.Println("秋季")
	} else if x >= 10 && x <= 12 {
		fmt.Println("冬季")
	} else {
		fmt.Println("没有这个月份")
	}
}

func SeasonSwitch(x int) {
	switch {
	case x >= 1 && x <= 3:
		fmt.Println("春季")
	case x >= 4 && x <= 6:
		fmt.Println("夏季")
	case x >= 7 && x <= 9:
		fmt.Println("秋季")
	case x >= 10 && x <= 12:
		fmt.Println("冬季")
	default:
		fmt.Println("没有这个月份")
	}
}

// 3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
type Student struct {
	Name        string
	Language    int
	Mathematics int
	English     int
}

func Class(x int) {
	var class = []Student{}
	// 创建一个班级的同学
	for i := 1; i <= x; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		class = append(class, Student{Name: fmt.Sprintf("stu_%d", i), Language: rand.Intn(100), Mathematics: rand.Intn(100), English: rand.Intn(100)})
	}

	var score int
	var manTime int

	for i := 0; i < len(class); i++ {
		stuScore := class[i].English + class[i].Language + class[i].Mathematics
		stuAvg := stuScore / 3
		fmt.Printf("同学: %s 的平均分为: %d\n", class[i].Name, stuAvg)

		if stuAvg < 60 {
			manTime++
		}

		score += stuScore
	}

	fmt.Printf("班级平均分为: %d\n", score/x)
	fmt.Printf("全班低于60分的同学人数有: %d 人\n", manTime)
}

func main() {
	// 1. 随机初始化两个8*5的矩阵，求两个矩阵的和(逐元素相加)
	arr1 := matrix_create()
	arr2 := matrix_create()

	fmt.Println(arr1)
	fmt.Println(arr2)

	arr := SumMatrix(arr1, arr2)
	fmt.Println(arr)

	// 2.给定月份，判断属于哪个季节。分别用if和switch实现
	SeasonIf(5)
	SeasonSwitch(10)

	// 3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
	Class(10)
}
