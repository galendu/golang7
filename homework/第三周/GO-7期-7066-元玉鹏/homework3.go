package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
*/

// Generate a Slice which cap == rows *cols and every element is a random number
func generateInteger(cap int) func() []int {
	return func() []int {
		randSlice := make([]int, 0, cap)
		rand.Seed(time.Now().UnixNano()) // avoid to generate the same random number every time
		for i := 0; i < cap; i++ {
			randSlice = append(randSlice, rand.Intn(1000))
		}
		return randSlice
	}
}

// This functions is used to generater two sub slice by the given random slice
func NewSubSlice(rows, cols int) ([][]int, [][]int) {
	base := generateInteger(rows * cols)
	baseSlice1, baseSlice2 := base(), base()
	subSlice1, subSlice2 := make([][]int, 0, rows), make([][]int, 0, rows)

	for i := 0; i < rows; i++ {
		subSlice1 = append(subSlice1, baseSlice1[i*cols:cols*(i+1)])
		subSlice2 = append(subSlice2, baseSlice2[i*cols:cols*(i+1)])
	}
	return subSlice1, subSlice2
}

// Question one
func q1(x, y [][]int, sum *([8][5]int)) [8][5]int {

	//calculate every element of the two given slice
	for row := 0; row < len(x); row++ {
		for col := 0; col < len(x[0]); col++ {
			sum[row][col] = x[row][col] + y[row][col]
		}
	}

	fmt.Println("Slice1 is:", x)
	fmt.Println("Slice1 is:", y)
	fmt.Println("Result is:", *sum)
	return *sum
}

/*
2.给定月份，判断属于哪个季节。分别用if和switch实现
*/

func q2(month string) (season string) {
	switch month {
	case "Jan", "Feb", "Mar", "January", "February", "March":
		return "spring"
	case "Apr", "May", "Jun", "April", "June":
		return "summer"
	case "Jul", "Aug", "Sep", "July", "August", "September":
		return "autumn"
	case "Oct", "Nov", "Dec", "October", "November", "December":
		return "winter"
	default:
		return "unknown"
	}
}

/*
3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
*/

// Define a student struct which contains three goals and name
type Student struct {
	Name        string
	Mathematics int
	Chinese     int
	English     int
}

// define a function for Student struct which is used to calculate the student average scole
func (s Student) Average() float32 {
	return (float32(s.Mathematics) + float32(s.English) + float32(s.Chinese)) / 3
}

type MyType []Student

// to calculate all the student's average math
func (m MyType) SumMathAvg() float32 {
	var sum int
	for _, student := range m {
		sum += student.Mathematics
	}
	return float32(sum) / float32(len(m))
}

// to calculate all the student's average chinese
func (m MyType) SumChineseAvg() float32 {
	var sum int
	for _, student := range m {
		sum += student.Chinese
	}
	return float32(sum) / float32(len(m))
}

// to calculate all the student's average english
func (m MyType) SumEnglishAvg() float32 {
	var sum int
	for _, student := range m {
		sum += student.English
	}
	return float32(sum) / float32(len(m))
}

// to retuen a slice containing all the student's average scores'
func (m MyType) SubAvg() map[string]float32 {
	ret := make(map[string]float32, len(m))
	for _, student := range m {
		ret[student.Name] = student.Average()
	}
	return ret
}

// calculate the number of failed in the exam
func (m MyType) SumFailed() int {
	ret := 0
	for _, student := range m {
		if student.Average() < float32(60) {
			ret++
		}
	}
	return ret
}

// generate a slice containing the student
func NewStudentSlice(num int) MyType {
	m := make(MyType, 0, num)
	for i := 0; i < num; i++ {
		rand.Seed(time.Now().UnixNano())
		m = append(m, Student{
			Name:        fmt.Sprintf("name%d", i),
			Mathematics: rand.Intn(100),
			Chinese:     rand.Intn(100),
			English:     rand.Intn(100),
		})
	}
	return m
}

func main() {
	// q1
	fmt.Println("----------第一题----------")
	slc1, slc2 := NewSubSlice(8, 5)
	q1(slc1, slc2, &([8][5]int{}))

	//q2
	fmt.Println("----------第二题----------")
	fmt.Println("May is", q2("May"))

	//q3
	fmt.Println("----------第三题----------")
	class := NewStudentSlice(5)
	fmt.Printf("班级所有同学的情况为:\n%v\n", class)
	fmt.Printf("每位同学的平均分为:\n%v\n", class.SubAvg())
	fmt.Printf("全部同学的数学平均分为:%v\n", class.SumMathAvg())
	fmt.Printf("全部同学的语文平均分为:%v\n", class.SumChineseAvg())
	fmt.Printf("全部同学的英语平均分为:%v\n", class.SumEnglishAvg())
	fmt.Printf("全部同学中不及格的人数为:%v\n", class.SumFailed())

}
