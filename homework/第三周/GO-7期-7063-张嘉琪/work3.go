package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*----------------------作业1------------------------------
1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
*/

type matrix struct {
	s *[8][5]int
}

//初始化matrix类型
func (m matrix) init() {

	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			(*(m.s))[i][j] = rand.Intn(128)
		}
	}
}

//matrix相加
func (m matrix) add(m1 matrix) {

	sum := [8][8]int{}

	//矩阵乘法：sum[i][j]=a[i][]*b[][j]
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			for k := 0; k < 5; k++ {
				sum[i][j] += (*(m.s))[i][k] * (*(m1.s))[j][k]
			}
			//打印sum[i][j]
			fmt.Printf("[%5d] ", sum[i][j])
		}
		//打印完一行，换行
		fmt.Println()
	}

}

/*-----------------------------------作业2------------------------------------
2.给定月份，判断属于哪个季节。分别用if和switch实现
*/

//输入月份，类型为int，
//用if判断是哪个季节
//打印判断结果
func season_if(a int) {
	if a > 0 && a < 4 {
		fmt.Printf("%d月份是春季\n", a)
	} else if a > 3 && a < 7 {
		fmt.Printf("%d月份是夏季\n", a)
	} else if a > 6 && a < 10 {
		fmt.Printf("%d月份是秋季\n", a)
	} else if a > 9 && a < 13 {
		fmt.Printf("%d月份是冬季\n", a)
	} else {
		fmt.Println("输入正确的月份")
	}
}

//输入月份，类型为int，
//用switch判断是哪个季节
//打印判断结果
func season_switch(a int) {
	switch {
	case a > 0 && a < 4:
		fmt.Printf("%d月份是春季\n", a)
	case a > 3 && a < 7:
		fmt.Printf("%d月份是夏季\n", a)
	case a > 6 && a < 10:
		fmt.Printf("%d月份是秋季\n", a)
	case a > 9 && a < 13:
		fmt.Printf("%d月份是冬季\n", a)
	default:
		fmt.Println("输入正确的月份")
	}
}

/*----------------------作业3----------------------------
3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
*/

type student struct {
	name    string
	chinese int
	math    int
	english int
}

func getStudentAvg(stu []student) map[string]int {
	s := make(map[string]int)
	for _, v := range stu {
		s[v.name] = (v.chinese + v.english + v.math) / 3
	}

	return s
}

func getClassAvg(stu []student) map[string]int {
	cv := make(map[string]int)
	var sumc, summ, sume int
	for _, v := range stu {
		sumc += v.chinese
		summ += v.math
		sume += v.english
	}

	cv["chinese"] = sumc / len(stu)
	cv["math"] = summ / len(stu)
	cv["english"] = sume / len(stu)

	return cv

}

func getFailStudent(sv map[string]int) int {
	var sumf int
	for _, v := range sv {
		if v < 60 {
			sumf += 1
		}
	}

	return sumf
}

/*--------------------------------------------------------------*/

func main() {

	/*---------------矩阵--------------------*/
	rand.Seed(time.Now().UnixNano())
	//用默认值初始化两个[8][5]int数组
	m1 := matrix{s: &[8][5]int{}}
	m2 := matrix{s: &[8][5]int{}}

	//分别调用初始化函数
	m1.init()
	m2.init()

	//m1*m2
	fmt.Println("m1*m2")
	m1.add(m2)
	fmt.Println()
	/*--------------月份-------------------------*/
	season_if(13)
	season_if(0)
	season_if(4)
	season_if(3)
	season_if(7)
	season_if(10)
	season_switch(13)
	season_switch(0)
	season_switch(4)
	season_switch(3)
	season_switch(7)
	season_switch(10)

	/*----------------平均分--------------------*/
	//每个学生的平均分
	ss := []student{{"zs", 98, 79, 100}, {"ls", 99, 76, 57}, {"ww", 46, 26, 74}}
	sv := getStudentAvg(ss)
	fmt.Println(sv)

	//班级的三门课程的平均分
	cv := getClassAvg(ss)
	fmt.Println(cv)

	//统计不及格学生的数量
	fmt.Printf("不及格学生数为:%d", getFailStudent(sv))
}
