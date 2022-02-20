package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//fcun1()
	//func2(7)
	func3()
}

//1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
func fcun1() {
	const (
		w = 8
		h = 5
	)
	A := [w][h]int{}
	for m := 0; m < w; m++ {
		for n := 0; n < h; n++ {
			A[m][n] = rand.Intn(20)
		}
	}
	B := [w][h]int{}
	for m := 0; m < w; m++ {
		for n := 0; n < h; n++ {
			B[m][n] = rand.Intn(20)
		}
	}
	C := [w][h]int{}
	sum := 0
	for m := 0; m < w; m++ {
		for n := 0; n < h; n++ {
			C[m][n] = A[m][n] + B[m][n]
			//fmt.Printf("%d\n",C)
			sum += C[m][n]
		}
	}
	fmt.Printf("最终sum=%d\n", sum)
}

//2.给定月份，判断属于哪个季节。分别用if和switch实现
func func2(num int) {
	switch {
	case num >= 1 && num <= 3:
		fmt.Printf("%d是春季", num)
	case num >= 4 && num <= 6:
		fmt.Printf("%d是夏季", num)
	case num >= 7 && num <= 9:
		fmt.Printf("%d是秋季", num)
	case num >= 10 && num <= 12:
		fmt.Printf("%d是冬季", num)
	}
}

//3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
type class struct {
	name                                               string
	ChineseScore, MethodScore, EnglishScore, Evelement int
}

func func3() {
	//定义学生列表
	nameList := []string{"小红", "小张", "小李", "小白", "小宋"}
	//定义一个class struct的切片
	classList := []class{}
	//平均分未到达60分的人数
	sum := 0
	//语文全班总分数
	Ch := 0
	//数学全班总分数
	Sh := 0
	//英语全班总分数
	Yi := 0

	for i := 0; i < len(nameList); i++ {
		name := nameList[i]
		ChineseScore := rand.Intn(100)
		MethodScore := rand.Intn(100)
		EnglishScore := rand.Intn(100)
		//计算平均分
		Evelement := (ChineseScore + MethodScore + EnglishScore) / 3
		if Evelement < 60 {
			sum += 1
		}
		//计算全班语文成绩总数
		Ch += ChineseScore
		//计算全班数学成绩总数
		Sh += MethodScore
		//计算全班英语成绩总数
		Yi += EnglishScore
		// 把语文，数学，英语，平均分初始化并添加到切片里
		classList = append(classList, class{name, ChineseScore, MethodScore, EnglishScore, Evelement})
		fmt.Printf("%d的语文成绩是%d,数学成绩是%d,英语成绩是%d,平均成绩是%d\n", name, ChineseScore, MethodScore, EnglishScore, Evelement)
	}
	fmt.Printf("共有%d位同学平均分没到达60分,语文平均分是%d,数学平均分是%d,英语平均分是%d\n", sum, Ch/len(nameList), Sh/len(nameList), Yi/len(nameList))
}
