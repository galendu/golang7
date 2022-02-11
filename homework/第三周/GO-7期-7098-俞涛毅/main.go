package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func get_sum() {
	sum := 0
	a := [8][5]int{}
	b := [8][5]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			a[i][j] = rand.Intn(1000)
			b[i][j] = rand.Intn(1000)
			sum += a[i][j] + b[i][j]

		}
	}
	fmt.Printf("两个矩阵求和结果为: %d", sum)
}

func season_if(m int) {
	if m >= 2 && m <= 4 {
		fmt.Printf("%d月当前季节是春季\n", m)
	} else if m >= 5 && m <= 7 {
		fmt.Printf("%d月当前季节是夏季\n", m)
	} else if m >= 8 && m <= 10 {
		fmt.Printf("%d月当前季节是秋季\n", m)
	} else {
		fmt.Printf("%d月当前季节是冬季\n", m)
	}

}

func season_sw(m int) {
	switch {
	case m >= 2 && m <= 4:
		fmt.Printf("%d月当前季节是春季\n", m)
	case m >= 5 && m <= 7:
		fmt.Printf("%d月当前季节是夏季\n", m)
	case m >= 8 && m <= 10:
		fmt.Printf("%d月当前季节是秋季\n", m)
	default:
		fmt.Printf("%d月当前季节是冬季\n", m)

	}
}

type student struct {
	name                 string
	yuwen, shuxue, waiyu int
}

func get_all_stu_avg(num int) {
	
	stu := student{}
	s := make([]student, 0, num)
	sum_y := 0
	sum_w := 0
	sum_s := 0
	avg_60 := 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		stu.yuwen = rand.Intn(100) + 1
		stu.shuxue = rand.Intn(100) + 1
		stu.waiyu = rand.Intn(100) + 1
		stu.name = strconv.Itoa(i)
		s = append(s, stu)
		avg := (s[i].yuwen + s[i].shuxue + s[i].waiyu) / 3
		fmt.Printf("学生%2d 成绩： 语文 %3d 数学 %3d 外语 %3d  平均分：%3d \n", i, s[i].yuwen, s[i].shuxue, s[i].waiyu, avg)
		sum_y += s[i].yuwen
		sum_s += s[i].shuxue
		sum_w += s[i].waiyu
		if avg < 60 {
			avg_60 += 1
		}

	}
	fmt.Printf("语文的平均分是  %5d\n", sum_y/num)
	fmt.Printf("数学的平均分是  %5d\n", sum_s/num)
	fmt.Printf("外语的平均分是  %5d\n", sum_w/num)
	fmt.Printf("全班低于平均分60的同学有%d个\n", avg_60)

}

func main() {

	// 1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
	get_sum()
	// 2.给定月份，判断属于哪个季节。分别用if和switch实现
	season_if(5)
	season_sw(12)
	// 3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
	get_all_stu_avg(50)
}

// 第一题可以实现生成矩阵的方法函数
// 第二题思考下是不是少了一个逻辑判断
