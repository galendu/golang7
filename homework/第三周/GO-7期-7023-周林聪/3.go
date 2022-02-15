package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
func one3() {
	arr := [8][5]int{}
	brr := [8][5]int{}
	crr := [8][5]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			arr[i][j] = rand.Intn(66)
			brr[i][j] = rand.Intn(88)
			crr[i][j] = arr[i][j] + brr[i][j]
		}
	}
	//fmt.Printf("arr=%3v\n", arr)
	//fmt.Printf("brr=%3v\n", brr)
	fmt.Printf("crr=%3v\n", crr)
}

//2.给定月份，判断属于哪个季节。分别用if和switch实现
func two3if(month string) {
	//用if实现
	if month == "1月" || month == "2月" || month == "3月" {
		fmt.Printf("%s是春季\n", month)
	} else if month == "4月" || month == "5月" || month == "6月" {
		fmt.Printf("%s是夏季\n", month)
	} else if month == "7月" || month == "8月" || month == "9月" {
		fmt.Printf("%s是秋季\n", month)
	} else if month == "10月" || month == "11月" || month == "12月" {
		fmt.Printf("%s是冬季\n", month)
	} else {
		fmt.Println("请输入正确的格式，如：1月")
	}
}

func two3switch(month string) {
	//用switch实现
	switch month {
	case "1月", "2月", "3月":
		fmt.Printf("%s是春季\n", month)
	case "4月", "5月", "6月":
		fmt.Printf("%s是夏季\n", month)
	case "7月", "8月", "9月":
		fmt.Printf("%s是秋季\n", month)
	case "10月", "11月", "12月":
		fmt.Printf("%s是秋季\n", month)
	default:
		fmt.Println("请输入正确的格式，如：1月")
	}
}

//3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
type student struct {
	name                 string
	yuwen, shuxue, waiyu float32
}

func (aa student) namefangfa(i int) string {
	aa.name = "学生"
	return fmt.Sprint(aa.name + strconv.Itoa(i))
}

var (
	yuwensum, shuxuesum, waiyusum, stuAve float32
	bujigeNum                             int
)

func three3() {
	arr3 := []student{}
	for i := 0; i < 10; i++ {
		a := student{}
		a.name = a.namefangfa(i + 1)
		a.yuwen = float32(rand.Intn(60) + 40)
		a.shuxue = float32(rand.Intn(60) + 40)
		a.waiyu = float32(rand.Intn(60) + 40)
		arr3 = append(arr3, a)

		stuscore := a.yuwen + a.shuxue + a.waiyu
		stuAve = stuscore / 3
		fmt.Printf("学生%d的平均分是%0.1f\n", i+1, stuAve)
		if stuAve < 60 {
			bujigeNum += 1
		}

		yuwensum += a.yuwen
		shuxuesum += a.shuxue
		waiyusum += a.waiyu
	}
	//fmt.Println(arr3)

	yuwenAve := yuwensum / float32(len(arr3))
	shuxueAve := shuxuesum / float32(len(arr3))
	waiyuAve := waiyusum / float32(len(arr3))

	fmt.Printf("全班语文平均分：%0.1f 全班数学平均分：%0.1f 全班外语平均分：%0.1f\n", yuwenAve, shuxueAve, waiyuAve)
	fmt.Printf("全班平均分低于60分的人数有%d个", bujigeNum)

}

func main() {
	//第一题
	one3()

	//第二题
	two3if("12月")
	two3switch("6月")

	//第三题
	three3()
}

// 第一个可以实现一个生成矩阵的方法 其它的可以优化下逻辑
