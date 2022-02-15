package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func sum() {
	//创建矩阵;好像不能创建未知长度的矩阵
	var a = [8][5]int{}
	var b = [8][5]int{}
	var total = 0
	rand.Seed(time.Now().Unix())
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			a[i][j] = rand.Intn(100)
			b[i][j] = rand.Intn(100)
		}
	}

	//输出举矩阵
	fmt.Println("a的矩阵输出：\n", a)
	fmt.Println("b的矩阵输出：\n", b)

	//输出矩阵和数据
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			total += a[i][j]
			total += b[i][j]
		}
	}
	fmt.Printf("矩阵和数据：%d\n", total)

}

func checkSeasonSwitch() {
	//获取当前月份
	var season int
	season = int(time.Now().Month())
	switch {
	case season <= 3:
		fmt.Printf("if: it is spring \n")
	case season <= 6:
		fmt.Printf("if: it is summer \n")
	case season <= 9:
		fmt.Printf("if: it is autumn \n")
	case season <= 12:
		fmt.Printf("if: it is winter \n")
	}
}

func checkSeasonIf() {
	//获取当前月份
	var season int
	season = int(time.Now().Month())
	if season <= 3 {
		fmt.Printf("switch: it is spring \n")
	} else if season <= 6 && season > 3 {
		fmt.Printf("switch: it is summer \n")
	} else if season <= 9 && season > 6 {
		fmt.Printf("switch: it is autumn \n")
	} else {
		fmt.Printf("switch: it is winter \n")
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
	//1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
	sum()
	//2.给定月份，判断属于哪个季节。分别用if和switch实现
	checkSeasonSwitch()
	checkSeasonIf()
	//3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
	average(5)
}

// 逻辑是正确的。不过思考下第二题是不是遗漏逻辑
