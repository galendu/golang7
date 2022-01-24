package main

import (
	"fmt"
	"math/rand"
)

// 1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）

func nets_for() {

	A := [8][5]int{}
	//1.1初始化二维数组
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			A[i][j] = rand.Intn(100)
		}
	}
	// fmt.Println(A)

	B := [8][5]int{}
	//1.2初始化二维数组
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			B[i][j] = rand.Intn(100)

		}

	}
	// fmt.Println(B)
	//1.3 求和
	sum := [8][5]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			sum[i][j] = B[i][j] + A[i][j]

		}

	}
	fmt.Println(sum)
}

// 2.给定月份，判断属于哪个季节。分别用if和switch实现
//1.用if判断
func if_month(month int) {
	if month >= 1 && month <= 3 {
		fmt.Printf("%d月是春天\n", month)
	} else if month >= 4 && month <= 6 {
		fmt.Printf("%d月是夏天\n", month)
	} else if month >= 7 && month <= 9 {
		fmt.Printf("%d月是秋天\n", month)
	} else if month >= 10 && month <= 12 {
		fmt.Printf("%d月是冬天\n", month)
	} else {
		fmt.Printf("四季中没有%d这个月份\n", month)
	}
}

//2.用switch判断
func switch_month(month int) {
	switch {
	case month >= 1 && month <= 3:
		fmt.Printf("%d月是春天\n", month)
	case month >= 4 && month <= 6:
		fmt.Printf("%d月是夏天\n", month)
	case month >= 7 && month <= 9:
		fmt.Printf("%d月是秋天\n", month)
	case month >= 9 && month <= 12:
		fmt.Printf("%d月是冬天\n", month)
	default:
		fmt.Printf("四季中没有%d这个月份\n", month)
	}

}

// 3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
//3.1 创建结构体
type student struct {
	name    string
	English int
	math    int
	chinese int
}

func num_ang() {
	//3.2 用一个slice容纳一个班的同学
	//初始化班级里面同学的成绩
	class := []student{}
	// Name := [...]int{1, 2, 3}
	name := [...]string{"小明", "小白", "小张", "小李", "小马", "小黄", "小李", "小钉", "小黑"}
	for i := 0; i < len(name); i++ {
		English := rand.Intn(100)
		math := rand.Intn(100)
		chinese := rand.Intn(100)
		class = append(class, student{string(name[i]), English, math, chinese})

	}
	// fmt.Println(class)
	//3.3 每位同学的平均分
	English_sum := 0
	math_sum := 0
	chinese_sum := 0
	count := 0
	for i := 0; i < len(class); i++ {
		avg := (class[i].math + class[i].English + class[i].chinese) / 3
		fmt.Printf("%s同学的平均分是%d\n", class[i].name, avg)
		//3.4全班同学平均分低于60的有几位
		if avg < 60 {
			count++
		}
		//3.5和整个班三门课的平均分
		English_sum += class[i].English
		math_sum += class[i].math
		chinese_sum += class[i].chinese
	}
	//3.4全班同学平均分低于60的有几位
	fmt.Printf("全班同学平均分低于60的有%d位同学\n", count)
	//3.5和整个班三门课的平均分
	fmt.Printf("整个班English课的平均分是%d\n", English_sum/(len(name)))
	fmt.Printf("整个班math课的平均分是%d\n", math_sum/(len(name)))
	fmt.Printf("整个班chinese课的平均分是%d\n", chinese_sum/(len(name)))

}

//主函数
func main() {
	nets_for()
	if_month(5)
	switch_month(12)
	num_ang()

}
