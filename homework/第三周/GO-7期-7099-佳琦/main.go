// 1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
// 2.给定月份，判断属于哪个季节。分别用if和switch实现
// 3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位

package main

import (
	"fmt"
	"math"
	"math/rand"
)

// 1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
type Matrix struct {
	value [][]int
	row int
	column int
}

func (m *Matrix) add(m1 *Matrix) *Matrix{
	maxRow := int(math.Max(float64(m.row), float64(m1.row)))
	maxColumn := int(math.Max(float64(m.column), float64(m1.column)))
	resM := generateMatrixWithFunc(maxRow, maxColumn, func() int {
		return 0
	})
	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxColumn; j++ {
			resM.value[i][j] = m.value[i][j] + m1.value[i][j]
		}
	}
	return resM
}

func generateMatrixWithFunc(row, column int, genFn func() int) *Matrix {
	arr := make([][]int, row)
	for idx, _ := range(arr) {
		arr[idx] = make([]int, column)
		for cIdx, _ := range(arr[idx]) {
			arr[idx][cIdx] = genFn()
		}
	}
	m := &Matrix{
		value: arr,
		row: row,
		column: column,
	}
	return m
}

func generateMatrix(row, column int) *Matrix {
	return generateMatrixWithFunc(row, column, func() int {
		return rand.Intn(32)
	})
}

// 2.给定月份，判断属于哪个季节。分别用if和switch实现
func getSeasonByMonth(m int) string {
	if m < 0 || m > 12 {
		return "err"
	} else if m < 4 {
		return "spring"
	} else if m < 7 {
		return "summer"
	} else if m < 10 {
		return "autumn"
	} else {
		return "winter"
	}
}

func getSeasonByMonth1(m int) string {
	var season string
	switch true {
	case m > 0 && m < 4:
		season = "spring"
	case m > 3 && m < 7:
		season = "summer"
	case m > 6 && m < 10:
		season = "autumn"
	case m > 9 && m <= 12:
		season = "winter"
	}
	return season
}


// 3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位

type student struct {
	name string
	chineseScore float32
	mathScore float32
	englishScore float32
}

func (s *student) avg () float32 {
	return (s.chineseScore + s.mathScore + s.englishScore) / 3
}

func (s *student) mockScore (prop string) float32 {
	halfList := [2]float32{0, 0.5}
	score := float32(rand.Intn(100)) + halfList[rand.Intn(2)]
	switch prop {
	case "chineseScore":
		s.chineseScore = score
	case "mathScore":
		s.mathScore = score
	case "englishScore":
		s.englishScore = score
	}
	return score
}

type class struct {
	students []student
}

func (c *class) entrance(n uint8) *class{
	var students []student = make([]student, n)
	for idx, stu := range(students) {
		stu.mockScore("chineseScore")
		stu.mockScore("mathScore")
		stu.mockScore("englishScore")
		students[idx] = stu
	}
	c.students = students
	return c
}

func (c class) everyoneAvg() []float32{
	var avgList []float32
	for _, stu := range(c.students) {
		avgList = append(avgList, stu.avg())
	}
	return avgList
}

func (c class) classAvg() float32 {
	avgList := c.everyoneAvg()
	var sum float32
	for _, v := range avgList {
		sum += v
	}
	return sum / float32(len(avgList))
}

func (c class) passCount() int {
	avgList := c.everyoneAvg()
	var count int
	for _, val := range avgList {
		if val >= 60 {
			count++
		}
	}
	return count
}

func (c class) chineseScore() float32 {
	students := c.students
	var sum float32
	for _, stu := range students {
		sum += stu.chineseScore
	}
	return sum / float32(len(students))
}

func (c class) mathScore() float32 {
	students := c.students
	var sum float32
	for _, stu := range students {
		sum += stu.mathScore
	}
	return sum / float32(len(students))
}

func (c class) englishScore() float32 {
	students := c.students
	var sum float32
	for _, stu := range students {
		sum += stu.englishScore
	}
	return sum / float32(len(students))
}

func main() {
	// 1
	m1 := generateMatrix(8, 5)
	m2 := generateMatrix(8, 5)
	fmt.Println("矩阵1: ", m1, "\n矩阵2: ", m2)
	resM := m1.add(m2)
	fmt.Println("结果: ", resM)
	// 2
	fmt.Println("if季节: ", getSeasonByMonth(4))
	fmt.Println("switch季节: ", getSeasonByMonth1(4))
	// 3
	c := class{}
	c.entrance(5)
	fmt.Println("学生列表: ", c.students)
	avgList := c.everyoneAvg()
	fmt.Println("每个人的平均分: ", avgList)
	fmt.Println("平均分及格人数: ", c.passCount())
	classAvg := c.classAvg()
	fmt.Println("班级平均分: ", classAvg)
	fmt.Println("语文平均分: ", c.chineseScore())
	fmt.Println("数学平均分: ", c.mathScore())
	fmt.Println("英语平均分: ", c.englishScore())
}