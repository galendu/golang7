package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// 1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）

func MakeRandomMatrix(length int, width int) [][]int {
	// 如果切片初始化长度和容量相等，则可以使用下标进行重新赋值
	matrix := make([][]int, length)
	for i := 0; i < length; i++ {
		s := make([]int, width)
		for j := 0; j < width; j++ {
			s[j] = rand.Intn(100)
		}
		matrix[i] = s
	}

	return matrix
}

func MatrixAdd(m1 [][]int, m2 [][]int) [][]int {
	// 如果切片初始化时指定了长度为0，只能使用append追加元素
	matrix := make([][]int, 0, len(m1))
	for i := 0; i < len(m1); i++ {
		s := make([]int, 0, len(m1[i]))
		for j := 0; j < len(m1[i]); j++ {
			s = append(s, m1[i][j]+m2[i][j])
		}
		matrix = append(matrix, s)
	}
	return matrix
}

func MatrixPrint(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Println()
	}
}

// 2.给定月份，判断属于哪个季节。分别用if和switch实现

func PrintSeason(month int) {
	if month > 0 && month <= 3 {
		fmt.Println("春")
	} else if month > 3 && month <= 6 {
		fmt.Println("夏")
	} else if month > 6 && month <= 9 {
		fmt.Println("秋")
	} else if month > 9 && month <= 12 {
		fmt.Println("冬")
	} else {
		fmt.Println("请输入 1~12 的数字")
	}
}

// 3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位

type Student struct {
	Name    string
	CnScore int
	EnScore int
	MtScore int
}

func InitStudentScore(num int) []Student {
	s := make([]Student, 0, num)
	for i := 0; i < num; i++ {
		u := Student{"name-" + strconv.Itoa(i), rand.Intn(100), rand.Intn(100), rand.Intn(100)}
		s = append(s, u)
	}
	return s
}

func (s Student) AvgScore() int {
	return (s.CnScore + s.EnScore + s.MtScore) / 3
}

func ClassAvgScore(s []Student) int {
	sum := 0
	for _, stu := range s {
		sum = sum + stu.CnScore + stu.EnScore + stu.MtScore
	}
	return sum / 3 / len(s)
}

func main() {

	// m1 := MakeRandomMatrix(8, 5)
	// m2 := MakeRandomMatrix(8, 5)
	// sum := MatrixAdd(m1, m2)
	// MatrixPrint(m1)
	// fmt.Println("--------------------------")
	// MatrixPrint(m2)
	// fmt.Println("--------------------------")
	// MatrixPrint(sum)

	// PrintSeason(0)
	// PrintSeason(1)
	// PrintSeason(10)
	// PrintSeason(15)

	s := InitStudentScore(100)
	fmt.Println(ClassAvgScore(s))

	for _, stu := range s {
		avg := stu.AvgScore()
		if avg < 60 {
			fmt.Printf("Student %s 平均分 %d 小于60分\n", stu.Name, avg)
		}

	}

}

// 第一题可以优化下
// 第二题少了switch方法
// 第三题未完成题目要求
