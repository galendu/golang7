package main

import (
	"fmt"
	"math/rand"
)

//1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
func jz() {

	//设定行、列常量
	const (
		row = 8
		col = 5
	)
	//初始化矩阵A
	A := [row][col]int{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			A[i][j] = rand.Intn(100)
		}
	}

	//初始化矩阵B
	B := [row][col]int{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			B[i][j] = rand.Intn(100)
		}
	}

	//初始化矩阵C
	C := [row][col]int{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			C[i][j] = A[i][j] + B[i][j] //求和
		}
	}

	//通过遍历显示矩阵(若直接打印C，显示为数组，不直观)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%6d", C[i][j]) // %6d使结果保持右对齐
		}
		fmt.Println() //打印空行，否则每行数据会挤在一起
	}
}

//2.给定月份，判断属于哪个季节。分别用if和switch实现
func season(month int) {
	if month > 0 && month < 4 {
		fmt.Printf("%d月是春季\n", month)
	} else if month > 3 && month < 7 {
		fmt.Printf("%d月是夏季\n", month)
	} else if month > 6 && month < 10 {
		fmt.Printf("%d月是秋季\n", month)
	} else if month > 9 && month < 13 {
		fmt.Printf("%d月是冬季\n", month)
	} else {
		fmt.Println("月份不正确")
	}
}

func season_switch(month int) {
	switch {
	case month > 0 && month < 4:
		fmt.Printf("%d月是春季\n", month)
	case month > 3 && month < 7:
		fmt.Printf("%d月是夏季\n", month)
	case month > 6 && month < 10:
		fmt.Printf("%d月是秋季\n", month)
	case month > 9 && month < 13:
		fmt.Printf("%d月是冬季\n", month)
	default:
		fmt.Println("月份不正确")
	}
}

//3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
type student struct {
	name    string
	chinese int
	math    int
	english int
}

func personal_score(s []student) {
	num := 0
	for _, v := range s {
		if (v.chinese+v.math+v.english)/3 < 60 {
			num += 1
		}
	}
	fmt.Printf("全班同学平均分低于60的有%d个\n", num)
}

func average_score(s []student) {
	sum_chinese := 0
	sum_math := 0
	sum_english := 0
	for _, v := range s {
		fmt.Printf("%s的科目平均分:%d\n", v.name, (v.chinese+v.math+v.english)/3)
		sum_chinese += v.chinese
		sum_math += v.math
		sum_english += v.english
	}
	fmt.Printf("全班语文平均分:%d\n", sum_chinese/len(s))
	fmt.Printf("全班数学平均分:%d\n", sum_math/len(s))
	fmt.Printf("全班英语平均分:%d\n", sum_english/len(s))

}

func main() {
	//answer1
	jz()
	//answer2
	season(12)
	season_switch(0)
	//answer3
	www := student{
		name:    "www",
		chinese: 60,
		math:    50,
		english: 40,
	}
	qqq := student{
		name:    "qqq",
		chinese: 40,
		math:    30,
		english: 20,
	}
	a := []student{www, qqq}
	average_score(a)
	personal_score(a)

}
