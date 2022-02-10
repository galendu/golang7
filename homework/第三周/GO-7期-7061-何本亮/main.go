package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
	1. 随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
*/

const (
	a = 8
	b = 5
)

func SumOfMatrix() {
	MatrixA := [a][b]float64{}
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			MatrixA[i][j] = rand.Float64()
		}
	}

	MatrixB := [a][b]float64{}
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			MatrixB[i][j] = rand.Float64()
		}
	}

	SumMatrix := [a][b]float64{}
	for i := 0; i < len(SumMatrix); i++ {
		for j := 0; j < len(SumMatrix[0]); j++ {
			sum := MatrixA[i][j] + MatrixB[i][j]
			SumMatrix[i][j] = sum
		}
	}

	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Printf("SumMatrix[%d][%d]=%f\n", i, j, SumMatrix[i][j])
		}
	}
}

/*
	2. 给定月份，判断属于哪个季节。分别用if和switch实现
*/
func JudgeSeasonWithIf(x int) {
	if x >= 1 && x <= 3 {
		fmt.Printf("%d月是春季\n", x)
	} else if x >= 4 && x <= 6 {
		fmt.Printf("%d月是夏季\n", x)
	} else if x >= 7 && x <= 9 {
		fmt.Printf("%d月是秋季\n", x)
	} else if x >= 10 && x <= 12 {
		fmt.Printf("%d月是冬季\n", x)
	} else {
		fmt.Printf("%d属于不合法的月份\n", x)
	}
}

func JudgeSeasonWithSwatch(x int) {
	switch x {
	case 1, 2, 3:
		{
			fmt.Printf("func %d月是春季\n", x)
		}
	case 4, 5, 6:
		{
			fmt.Printf("%d月是夏季\n", x)
		}
	case 7, 8, 9:
		{
			fmt.Printf("%d月是秋季\n", x)
		}
	case 10, 11, 12:
		{
			fmt.Printf("%d月是冬季\n", x)
		}
	default:
		fmt.Printf("%d属于不合法的月份\n", x)
	}

}

//创建学生数据结构体
type Student struct {
	Name         string
	ChineseScore int
	MathScore    int
	EnglishScore int
	average      float32
}

func average(n int) {
	//创建学科数据
	var stuSlice []Student
	//var stuMap = map[int]Student{}
	var stu Student
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		stu.Name = "student_" + strconv.Itoa(i+1)
		stu.ChineseScore = rand.Intn(100)
		stu.MathScore = rand.Intn(100)
		stu.EnglishScore = rand.Intn(100)
		fmt.Println("输出结构体数据：\n", stu)
		//if _,ok := stuMap[i];!ok{
		//	stuMap[i] = stu
		//}
		stuSlice = append(stuSlice, stu)
	}

	//输出结构体数据
	fmt.Println("输出切片结构体数据：", stuSlice)
	//fmt.Println("输出map结构体数据：",stuMap)

	//获取同学平均分和三科平均分
	var chineseAverageScore, mathAverageScore, englishAverageScore float32
	var chineseTotalScore, mathTotalScore, englishTotalScore = 0, 0, 0
	var peopleTotal = len(stuSlice)
	var notPass60 = 0
	for i := 0; i < len(stuSlice); i++ {
		chineseTotalScore += stuSlice[i].ChineseScore
		mathTotalScore += stuSlice[i].MathScore
		englishTotalScore += stuSlice[i].EnglishScore
		//获取平均分
		stuSlice[i].average = float32((stuSlice[i].ChineseScore + stuSlice[i].MathScore + stuSlice[i].EnglishScore) / 3)
		//获取平均分未及格人数
		if stuSlice[i].average < 60 {
			notPass60++
		}
	}
	//获取科目平均分
	chineseAverageScore = float32(chineseTotalScore / peopleTotal)
	mathAverageScore = float32(mathTotalScore / peopleTotal)
	englishAverageScore = float32(englishTotalScore / peopleTotal)

	fmt.Println("输出切片结构体数据(包含平均分)：", stuSlice)
	fmt.Println("平均分未达到60分同学数：", notPass60)
	fmt.Printf("Chinese:%.1f\n", chineseAverageScore)
	fmt.Printf("Math:%.1f\n", mathAverageScore)
	fmt.Printf("English:%.1f\n", englishAverageScore)

}

func main() {
	SumOfMatrix()

	x := 1
	JudgeSeasonWithIf(x)
	JudgeSeasonWithSwatch(x)

	average(5)
}

// 第一题可以实现一方方法函数生成矩阵
// 第二题代码switch实现方式需要复习下知识点
// 输出进行优化下
