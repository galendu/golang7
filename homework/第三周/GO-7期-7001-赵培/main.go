package main

import (
	"fmt"
	"math/rand"
	"time"
)

//问题1：随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）

func Matrix_make() [8][5]int {
	matrix := [8][5]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			num := rand.Intn(100) + 1
			matrix[i][j] = num
		}
	}
	return matrix

}

func Matrix_sum() {
	//构建一个时间种子
	rand.Seed(time.Now().UnixNano())

	//初始化两个矩阵
	matrix01 := Matrix_make()
	fmt.Println(matrix01)
	matrix02 := Matrix_make()
	fmt.Println(matrix02)

	sum := [8][5]int{}

	for i := 0; i < len(sum); i++ {
		for j := 0; j < len(sum[0]); j++ {
			sum[i][j] = matrix01[i][j] + matrix02[i][j]
		}
	}
	fmt.Println(sum)
}

//问题2：给定月份，判断属于哪个季节。分别用if和switch实现
func month_if() {
	fmt.Println("please input your month: ")
	var month int
	fmt.Scan(&month)
	if month >= 1 && month <= 3 {
		fmt.Println("The month you entered is [spring]")
	} else if month >= 4 && month <= 6 {
		fmt.Println("The month you entered is [summer]")
	} else if month >= 7 && month <= 9 {
		fmt.Println("The month you entered is [autumn]")
	} else if month >= 10 && month <= 12 {
		fmt.Println("Then month you entered is [winter]")
	} else {
		fmt.Println("you entered month [invalid]")
	}
}

func month_switch() {
	fmt.Println("please input your month: ")
	var month int
	fmt.Scan(&month)
	switch {
	case month >= 1 && month <= 3:
		fmt.Println("The month you entered is [spring]")
	case month >= 4 && month <= 6:
		fmt.Println("The month your entered is [summer]")
	case month >= 7 && month <= 9:
		fmt.Println("The month your entered is [autumn]")
	case month >= 10 && month <= 12:
		fmt.Println("The month your entered is [winter]")
	default:
		fmt.Println("The month your entered is [invalid]")
	}
}

//问题3：创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，
//求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
type student struct {
	Name        string
	Language    int
	Mathematics int
	English     int
}

func number_avg() {
	//构建一个种子
	rand.Seed(time.Now().UnixNano())

	//随机初始化一个班的同学的成绩
	stu := []student{}
	str := "abcdefghijabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(str); i++ {
		stu = append(stu, student{string(str[i]), rand.Intn(100) + 1, rand.Intn(100) + 1, rand.Intn(100) + 1})
	}

	// //定义一个map用来存储每位同学的平均分
	// avg01 := make(map[string]int)
	sum := 0
	count := 0
	for i := 0; i < len(stu); i++ {
		avg := (stu[i].Language + stu[i].Mathematics + stu[i].English) / 3
		fmt.Printf("%s同学的平均分为%d\n", stu[i].Name, avg)
		// avg01[stu[i].Name] = avg
		if avg < 60 {
			count++
		}
		sum += stu[i].Language + stu[i].Mathematics + stu[i].English
	}
	fmt.Printf("全班同学的平局成绩为%d\n", sum/len(stu))
	fmt.Printf("全班同学平均分低于60的有%d位\n", count)

}

func main() {
	Matrix_sum()
	month_if()
	month_switch()
	number_avg()
}

// 整体写的不错。第一个可以尝试下for循环不使用matrix[0]和sum[0]有没有其它方法