package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
*/
func matrixSummation() {
	const (
		LONG  = 8
		WIDTH = 5
	)
	A := [LONG][WIDTH]int{}
	B := [LONG][WIDTH]int{}
	rect := [LONG][WIDTH]int{}

	for i := 0; i < LONG; i++ {
		for j := 0; j < WIDTH; j++ {
			A[i][j] = rand.Intn(1000)
			B[i][j] = rand.Intn(1000)
			rect[i][j] = A[i][j] + B[i][j]
		}
	}
	fmt.Println(A)
	fmt.Println()
	fmt.Println(B)
	fmt.Println()
	fmt.Println(rect)
}

/*
2.给定月份，判断属于哪个季节。分别用if和switch实现
*/
func ifSeason(mouth int) {
	if mouth > 12 || mouth < 0 {
		fmt.Printf("请输入正确的月份")
	} else {
		if mouth > 6 && mouth <= 12 {
			if mouth > 9 && mouth <= 12 {
				fmt.Printf("%d月是冬天\n", mouth)
			} else {
				fmt.Printf("%d月是秋天\n", mouth)
			}
		} else {
			if mouth > 0 && mouth <= 3 {
				fmt.Printf("%d月是春天\n", mouth)
			} else {
				fmt.Printf("%d月是夏天\n", mouth)
			}
		}
	}
	fmt.Println()
}

func switchSeason(mouth int) {
	switch {
	case mouth > 12 || mouth < 0:
		fmt.Printf("请输入正确的月份")
	case mouth > 0 && mouth <= 3:
		fmt.Printf("%d月是春天\n", mouth)
	case mouth > 3 && mouth <= 6:
		fmt.Printf("%d月是夏天\n", mouth)
	case mouth > 6 && mouth <= 9:
		fmt.Printf("%d月是秋天\n", mouth)
	case mouth > 9 && mouth <= 12:
		fmt.Printf("%d月是冬天\n", mouth)
	}
	fmt.Println()
}

/*
3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
*/
type student struct {
	name                   string
	chinese, math, english float64
}

func (s student) average() float64 {
	average := s.chinese + s.math + s.english
	return average / float64(3)
}

func getClassAverage1(s int) { //s指定生成多少个学生
	var classAverage, chinese, math, english float64
	var sum int
	class := make([]student, s)

	for i := 0; i < s; i++ { //遍历slice
		a := student{}
		a.name = strconv.Itoa(i) + "a"
		rand.Seed(time.Now().UnixNano())
		a.chinese = float64(rand.Intn(80) + 20)
		a.math = float64(rand.Intn(80) + 20)
		a.english = float64(rand.Intn(80) + 20)
		class = append(class, a)

		fmt.Printf("%s的平均分是%g\n", a.name, a.average())
		if a.average() < 60 {
			sum = sum + 1
		}
		chinese += a.chinese
		math += a.math
		english += a.english
		classAverage += a.average()
	}

	fmt.Printf("全班语文平均分：%g,全班数学平均分：%g,全班英语平均分：%g,全班总分平均分：%g,平均分低于60分的有%d个\n", chinese/float64(s), math/float64(s), english/float64(s), classAverage/float64(s), sum)

}

func main() {
	matrixSummation()
	ifSeason(13)
	ifSeason(12)
	switchSeason(-1)
	switchSeason(5)

	getClassAverage1(20)

}
