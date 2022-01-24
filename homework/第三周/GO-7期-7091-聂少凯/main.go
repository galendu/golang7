package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
2.给定月份，判断属于哪个季节。分别用if和switch实现
3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位


*/

func getMatrix(line, colume int) [][]int {
	result := make([][]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < line; i++ {
		tmp := make([]int, 0)
		for j := 0; j < colume; j++ {
			tmp = append(tmp, rand.Intn(100))
		}
		result = append(result, tmp)
	}
	return result
}

func TwoMatrixAdd() [][]int {
	A := getMatrix(8, 5)
	B := getMatrix(8, 5)
	result := make([][]int, 0)
	for i := 0; i < len(A); i++ {
		tmp := make([]int, 0)
		for j := 0; j < len(A[0]); j++ {
			tmp = append(tmp, A[i][j]+B[i][j])
			fmt.Println(i, j, A[i][j]+B[i][j])
		}
		result = append(result, tmp)
	}
	fmt.Println(result)
	return A
}

func GetSeason(season int) {
	switch season {
	case 1, 2, 3:
		fmt.Println("春天")
	case 4, 5, 6:
		fmt.Println("夏天")
	case 7, 8, 9:
		fmt.Println("秋天")
	case 10, 11, 12:
		fmt.Println("冬天")
	}
}

type student struct {
	Name    string
	Chines  float64
	Math    float64
	English float64
}

func (s *student) GetAve() float64 {
	return (s.English + s.Math + s.Chines) / 3
}

type students []student

func (ss *students) GetAveForEveryone() {
	result := make(map[string]float64, 0)
	for _, stu := range *ss {
		result[stu.Name] = stu.GetAve()
	}
	fmt.Printf("全班共%v位同学,每一位的平均分是%v\n", len(*ss), result)
}

func (ss *students) GetAveForClass() {
	var result float64
	i := float64(1)
	for _, stu := range *ss {
		result += stu.GetAve()
		i++
	}
	fmt.Printf("全班共%v位同学,全班平均分是:%v\n", len(*ss), result/i)
}

func (ss *students) GetNmberScoreLow60() {
	result := 0
	for _, stu := range *ss {
		if stu.GetAve() < 60 {
			result++
		}
	}
	fmt.Printf("全班共%v位同学,低于60分有%v位\n", len(*ss), result)
}

func GetClass(num int) students {
	var class students
	const fullMark = 100
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= num; i++ {
		stu := student{}
		stu.Name = "同学" + strconv.Itoa(i)
		stu.Chines = float64(rand.Intn(fullMark))
		stu.Math = float64(rand.Intn(fullMark))
		stu.English = float64(rand.Intn(fullMark))
		class = append(class, stu)
	}
	return class
}
func Students(num int) {
	class := GetClass(num)
	class.GetAveForClass()
	class.GetAveForEveryone()
	class.GetNmberScoreLow60()
}
func main() {
	//1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
	//TwoMatrixAdd()

	//2.给定月份，判断属于哪个季节。分别用if和switch实现
	//GetSeason(8)

	//3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
	Students(50)

}
