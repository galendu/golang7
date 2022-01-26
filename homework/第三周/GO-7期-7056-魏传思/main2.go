package main

import (
	"fmt"
	"math/rand"
)

// 矩阵逐元素相加
func matriAdd(A [8][5]int, B [8][5]int) [8][5]int {
	sum := [8][5]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			sum[i][j] = A[i][j] + B[i][j]
		}
	}
	return sum
}

func jijie(month int) string {
	if month >= 1 && month <= 3 {
		return "春"
	} else if month >= 4 && month <= 6 {
		return "夏"
	} else if month >= 7 && month <= 9 {
		return "秋"
	} else if month >= 10 && month <= 12 {
		return "冬"
	} else {
		return fmt.Sprintf("非法的月份%d", month)
	}

}

func jijie2(month int) string {
	switch month {
	case 1, 2, 3:
		return "春"
	case 4, 5, 6:
		return "夏"
	case 7, 8, 9:
		return "秋"
	case 10, 11, 12:
		return "冬"
	default:
		return "不存在这个月份"
		// return fmt.Sprintf("非法的月份%d", month)
	}
}

type Student struct {
	Name         string
	ChineseScore float64
	MathScore    float64
	EnlishScore  float64
	AvgScore     float64
}

func (stu *Student) Average() {
	stu.AvgScore = (stu.ChineseScore + stu.MathScore + stu.EnlishScore) / 3
	fmt.Printf("aaa  %s %f\n", stu.Name, stu.AvgScore)
}

type Class struct {
	Students     []*Student
	ChineseScore float64
	MathScore    float64
	EnlishScore  float64
}

func (cls *Class) ChineseAvg() {
	if len(cls.Students) == 0 {
		return
	}
	var sum float64
	for _, stu := range cls.Students {
		sum += stu.ChineseScore
	}
	cls.ChineseScore = sum / float64(len(cls.Students))
}

func (cls *Class) MathAvg() {
	if len(cls.Students) == 0 {
		return
	}
	var sum float64
	for _, stu := range cls.Students {
		sum += stu.MathScore
	}
	cls.MathScore = sum / float64(len(cls.Students))
}

func (cls *Class) EnlishAvg() {
	if len(cls.Students) == 0 {
		return
	}
	var sum float64
	for _, stu := range cls.Students {
		sum += stu.EnlishScore
	}
	cls.EnlishScore = sum / float64(len(cls.Students))
}

func (cls *Class) Fail() int {
	n := 0
	for _, stu := range cls.Students {
		if stu.AvgScore < 60 {
			n++
		}

	}
	return n
}
func main() {
	A := [8][5]int{}
	B := [8][5]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			A[i][j] = rand.Intn(100)
			B[i][j] = rand.Intn(100)

		}
	}
	C := matriAdd(A, B)
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%5d", C[i][j])
		}
		fmt.Println()
	}

	month := 5
	fmt.Println(month, jijie(month))
	month = 13
	fmt.Println(month, jijie(month))
	month = 0
	fmt.Println(month, jijie(month))

	month = 13
	fmt.Println(month, jijie2(month))
	month = 0
	fmt.Println(month, jijie2(month))
	month = 1
	fmt.Println(month, jijie2(month))

	s1 := &Student{
		Name:         "小明",
		ChineseScore: 55,
		MathScore:    55,
		EnlishScore:  81,
	}
	s2 := &Student{
		Name:         "小华",
		ChineseScore: 45,
		MathScore:    45,
		EnlishScore:  60,
	}
	s3 := &Student{
		Name:         "小白",
		ChineseScore: 55,
		MathScore:    36,
		EnlishScore:  72,
	}
	class := Class{
		Students: []*Student{s1, s2, s3},
	}
	for _, stu := range class.Students {
		stu.Average()
	}
	class.ChineseAvg()
	class.MathAvg()
	class.EnlishAvg()
	fmt.Printf("不及格的同学有%d位\n", class.Fail())
}
