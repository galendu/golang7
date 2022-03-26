//1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
//2.给定月份，判断属于哪个季节。分别用if和switch实现
//3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func matrix() {
	const (
		ROW = 8
		COL = 5
	)

	A := [ROW][COL]int{}
	B := [ROW][COL]int{}
	C := [ROW][COL]int{}

	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			A[i][j] = rand.Intn(100)
			B[i][j] = rand.Intn(100)
			C[i][j] = A[i][j] + B[i][j]
			fmt.Printf("%3d ", C[i][j])
		}
		fmt.Println()
	}
}

func if_season(sea int) {
	//给定月份，判断属于哪个季节。分别用if和switch实现
	if sea < 1 || sea > 12 {
		fmt.Println("没有该月份")
		return
	}

	if sea <= 3 {
		fmt.Println("春季")
	} else if sea <= 6 {
		fmt.Println("夏季")
	} else if sea <= 9 {
		fmt.Println("秋季")
	} else if sea <= 12 {
		fmt.Println("冬季")
	}
}

func switch_season(sea int) {
	switch sea {
	case 1, 2, 3:
		fmt.Println("春季")
	case 4, 5, 6:
		fmt.Println("夏季")
	case 7, 8, 9:
		fmt.Println("秋季")
	case 10, 11, 12:
		fmt.Println("冬季")
	default:
		fmt.Println("没有该月份")
	}

}

/////////////

var num int = 0 // 低于60分的计数器
type Student struct {
	name                   string
	chinese, math, english int
}

type Class struct {
	Student
}

func (c Class) cls_avg(score int, person int) float64 {
	return float64(score / person)
}

func (c Class) below60(score float64) int {
	if score < 60 {
		num += 1
	}
	return num
}

func (s Student) stu_avg(chinese int, math int, english int) float64 {
	return float64((chinese + math + english) / 3)
}

func count_score(student_count int) {
	//创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
	var chinese, math, english int
	c_struct := Class{}
	class_slice := []Class{} // 班级容器
	below60 := 0             // 低于60分的个数

	// 填充一个班的信息
	for i := 1; i <= student_count; i++ {
		name := "student" + strconv.Itoa(i)
		chinese := rand.Intn(101)
		math := rand.Intn(101)
		english := rand.Intn(101)

		class_slice = append(class_slice, Class{Student{name: name, chinese: chinese, math: math, english: english}})
	}

	// 具体计算
	person_count := len(class_slice)
	for _, ele := range class_slice {
		avg_st := ele.stu_avg(ele.chinese, ele.math, ele.english)
		fmt.Printf("同学:%s, 平均分为: %.2f\n", ele.name, avg_st)
		below60 = ele.below60(avg_st)
		chinese += ele.chinese
		math += ele.math
		english += ele.english
	}

	fmt.Printf("全班平均分低于60分的同学有: %d名\n", below60)
	fmt.Printf("全班语文的平均分: %.2f\n", c_struct.cls_avg(chinese, person_count))
	fmt.Printf("全班数学的平均分: %.2f\n", c_struct.cls_avg(math, person_count))
	fmt.Printf("全班英语的平均分: %.2f\n", c_struct.cls_avg(english, person_count))
}

func main() {
	matrix()         // 第一题
	if_season(12)    //第二题
	switch_season(2) //第二题
	count_score(60)  // 第三题
}
