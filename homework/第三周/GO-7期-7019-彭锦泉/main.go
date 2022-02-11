package main

import (
	"fmt"
	"math/rand"
	"time"
)

//1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）

func First_task() {
	const X, Y int = 8, 5

	fmt.Println("A矩阵为：")
	A := [X][Y]float64{} //声明并初始化矩阵A
	rand.Seed(time.Now().Unix())
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			A[i][j] = rand.Float64()
			fmt.Printf("%f  ", A[i][j])
		}
		fmt.Println()
	}

	fmt.Println("B矩阵为：")
	B := [X][Y]float64{} //声明并初始化矩阵B
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			B[i][j] = rand.Float64()
			fmt.Printf("%f  ", B[i][j])
		}
		fmt.Println()
	}

	fmt.Println("两个矩阵的和为：")
	C := [X][Y]float64{} //声明矩阵C并计算以上两个矩阵的和 存放至矩阵C
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			C[i][j] = A[i][j] + B[i][j]
			fmt.Printf("%f  ", C[i][j])
		}
		fmt.Println()
	}
}

//2.给定月份，判断属于哪个季节。分别用if和switch实现

func Second_task(m uint8) string {
	if m >= 1 && m <= 12 {
		switch m {
		case 3, 4, 5:
			return "春天"
		case 6, 7, 8:
			return "夏天"
		case 9, 10, 11:
			return "秋天"
		case 12, 1, 2:
			return "冬天"
		}
	}
	return "无效的参数，请输入正确的月份。"
}

//3.创建一个student结构体，包含姓名和语数外三门课的成绩。
//用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位

type Student struct { //声明Student结构体
	Name        string
	Chinese     int
	Math        int
	Foreignlang int
}

//计算每位同学的平均分
func (u Student) stu_average() int {
	return (u.Chinese + u.Math + u.Foreignlang) / 3
}

//计算整个班级三门课的平均分
func subject_average(slice []Student, s string) int {
	if s == "Chinese" {
		sum := 0
		for _, ele := range slice {
			sum += ele.Chinese
		}
		return sum / len(slice)
	} else if s == "Math" {
		sum := 0
		for _, ele := range slice {
			sum += ele.Math
		}
		return sum / len(slice)
	} else if s == "Foreignlang" {
		sum := 0
		for _, ele := range slice {
			sum += ele.Foreignlang
		}
		return sum / len(slice)
	} else {
		goto L1
	}
L1:
	return -1
}

//计算全班平均分低于60的同学有多少个
func calculate(slice []Student) int {
	fail := 0
	for _, ele := range slice {
		if ele.stu_average() < 60 {
			fail += 1
		}
	}
	return fail
}

func Third_task() {
	class := make([]Student, 50)
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(class); i++ { //初始化每位同学的各科分数（随机）
		var u Student
		class[i] = u
		class[i].Name = "xxx"
		class[i].Chinese = rand.Intn(101)
		class[i].Math = rand.Intn(101)
		class[i].Foreignlang = rand.Intn(101)
	}
	fmt.Printf("这个班级一共有%d位学生\n", len(class))
	fmt.Println("每位同学三科平均分为：")
	for i := 0; i < len(class); i++ {
		fmt.Printf("%d ", class[i].stu_average())
		if (i+1)%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Printf("整个班级[语文]平均分：%d\n", subject_average(class, "Chinese"))
	fmt.Printf("整个班级[数学]平均分：%d\n", subject_average(class, "Math"))
	fmt.Printf("整个班级[外语]平均分：%d\n", subject_average(class, "Foreignlang"))
	fmt.Printf("全班中平均分低于60的同学有%d个\n", calculate(class))
}

func main() {
	First_task()
	fmt.Println("------------------------------------")
	var m uint8 = 8
	fmt.Printf("%d月份是%s\n", m, Second_task(m))
	fmt.Println("------------------------------------")
	Third_task()
}

// 输出打印的时候，要实现清晰明了
// 第二题是使用if和swich分别实现