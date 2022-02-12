package main

import (
	"fmt"
	"math/rand"
)

func season(Month int) string {
	if Month > 0 && Month <= 3 {
		return "春季"
	} else if Month > 3 && Month <= 6 {
		return "夏季"
	} else if Month > 6 && Month <= 9 {
		return "秋季"
	} else if Month > 9 && Month <= 12 {
		return "冬季"
	} else {
		return "输入有误无法正确判断当前季节..."
	}
}

func season1(Month int) string {
	switch Month {
	case 1, 2, 3:
		return "春季"
	case 4, 5, 6:
		return "夏季"
	case 7, 8, 9:
		return "秋季"
	case 10, 11, 12:
		return "冬季"
	default:
		return "输入有误无法正确判断当前季节..."
	}
}

//type socre interface {
//	Score() float32
//}

type Student struct {
	Name           string
	Yuwen          float32
	Shuxue         float32
	Yingyu         float32
	Numberofcourse int
}

type Class struct {
	Name     string
	Students []*Student
}

func (s Student) Score() float32 {

	return (s.Yingyu + s.Shuxue + s.Yuwen) / float32(s.Numberofcourse)
}

func (c Class) Score() (float32, float32, float32) {
	//if sum == 0 {
	//	fmt.Println("学生数不能为:0")
	//	return 0,0,0
	//}
	var students *Student
	var yunwen float32
	var shuxue float32
	var yingyu float32
	for _, students = range c.Students {
		yunwen += students.Yuwen
		shuxue += students.Shuxue
		yingyu += students.Yingyu
	}
	yuwensum := (yunwen) / float32(len(c.Students))
	shuxuesum := (shuxue) / float32(len(c.Students))
	yingyusum := (yingyu) / float32(len(c.Students))
	return yuwensum, shuxuesum, yingyusum
}

func (c Class) svg() int {
	var sum int

	for _, s := range c.Students {
		if s.Score() < 60 {
			sum++
		}
	}

	return sum
}

func matrix() [8][5]int64 {
	matrix := [8][5]int64{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			matrix[i][j] = int64(rand.Intn(128))
		}
	}
	return matrix
}

func matrixadd(m1, m2 [8][5]int64) [8][5]int64 {
	sum := [8][5]int64{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			sum[i][j] = m1[i][j] + m2[i][j]
		}
	}
	return sum
}

func main() {

	fmt.Printf("当前季节:%s\n", season(9))
	fmt.Printf("当前季节:%s\n", season1(13))

	score := &Student{
		Yuwen:          10,
		Shuxue:         10,
		Yingyu:         10,
		Numberofcourse: 3,
	}
	score1 := &Student{
		Yuwen:          100,
		Shuxue:         100,
		Yingyu:         100,
		Numberofcourse: 3,
	}

	score6 := &Student{
		Yuwen:          100,
		Shuxue:         100,
		Yingyu:         100,
		Numberofcourse: 3,
	}

	a := Student.Score(*score)
	fmt.Printf("班级个人平均分:%.2f\n", a)

	s := &Class{Name: "", Students: []*Student{score, score1, score6}}
	score2, score3, score4 := Class.Score(*s)
	fmt.Printf("班级语文平均分:%.2f 班级数学平均分:%.2f 班级英语平均分:%.2f\n", score2, score3, score4)

	score5 := Class.svg(*s)
	fmt.Printf("不及格的学生个数:%d\n", score5)

	fmt.Println(matrix())
	m1 := matrix()
	fmt.Println(m1)
	m2 := matrix()
	fmt.Println(m2)
	m3 := matrixadd(m1, m2)
	fmt.Println(m3)

}

// 完成的不错。第三题虽然实现了，但要注意题意
