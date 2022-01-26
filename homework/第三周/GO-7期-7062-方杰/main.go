package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

/*
1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
2.给定月份，判断属于哪个季节。分别用if和switch实现
3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
*/

const ROW, COL = 8, 5 // 定义矩阵行和列的数值

// 自定义分数结构体Score
type Score struct {
	Chinese float64
	Math    float64
	English float64
}

// 自定义学生结构体Student
type Student struct {
	Name     string
	Scores   Score // Score结构体类型
	AvgScore float64
}

// Student结构体方法，计算单个学生的平均分
func (stu Student) GenAvgScore() float64 {
	AvgScore := (stu.Scores.Chinese + stu.Scores.Math + stu.Scores.English) / 3
	return AvgScore
}

// 自定义班级结构体Classes
type Classes struct {
	Students    []Student // Student切片
	Chinese_avg float64
	Math_avg    float64
	English_avg float64
}

// Classes方法计算班级三门课平均分
func (cls Classes) GenAvg() []float64 {
	var chinese_sum, math_sum, english_sum float64
	cls_avgs := []float64{}
	nums := float64(len(cls.Students))

	for _, stu := range cls.Students {
		chinese_sum += stu.Scores.Chinese
		math_sum += stu.Scores.Math
		english_sum += stu.Scores.English
	}
	//chinese_avg := chinese_sum / 5
	//math_avg := math_sum / 5
	//english_avg := english_sum / 5

	//fmt.Println(len(cls.Students))
	//fmt.Printf("语文平均分:%.2f\n", chinese_avg)
	//fmt.Printf("数学平均分:%.2f\n", math_avg)
	//fmt.Printf("英语平均分:%.2f\n", english_avg)

	cls_avgs = append(cls_avgs, chinese_sum/nums)
	cls_avgs = append(cls_avgs, math_sum/nums)
	cls_avgs = append(cls_avgs, english_sum/nums)

	return cls_avgs
}

func GenArr() [ROW][COL]int {
	result := [ROW][COL]int{}
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			result[i][j] = rand.Intn(128)
		}
	}
	return result
}

func MatrixSum(arr1, arr2 [ROW][COL]int) [ROW][COL]int {
	result := [ROW][COL]int{}
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			result[i][j] = arr1[i][j] + arr2[i][j]
		}
	}

	return result
}

func Searson(month int) {
	if month >= 3 && month <= 5 {
		fmt.Printf("%d月是春天!\n", month)
	} else if month >= 6 && month <= 8 {
		fmt.Printf("%d月是夏天!\n", month)
	} else if month >= 9 && month <= 11 {
		fmt.Printf("%d月是秋天!\n", month)
	} else if month == 12 || (month >= 1 && month <= 2) {
		fmt.Printf("%d月是冬天!\n", month)
	} else {
		fmt.Println("月份不对，请传入正确的月份!")
	}

	switch month {
	case 3, 4, 5:
		fmt.Printf("%d月是春天!\n", month)
	case 6, 7, 8:
		fmt.Printf("%d月是夏天!\n", month)
	case 9, 10, 11:
		fmt.Printf("%d月是秋天!\n", month)
	case 1, 2, 12:
		fmt.Printf("%d月是冬天!\n", month)
	default:
		fmt.Println("月份不对，请传入正确的月份!")
	}
}

// 通过传入的最大最(max)小值(min)和数目(n)，生成一个切片，返回n个浮点数
func randFloats(min, max, num int) []float64 {
	var res []float64
	for i := 0; i < num; i++ {
		score := float64(min) + rand.Float64() + float64(rand.Intn(max-min))
		res = append(res, score)
	}
	return res
}

// 通过指定的数量(num)生成一个对应数量的学生类型切片
func GenStuScores(num int) []Student {
	students := []Student{}
	stu := Student{}
	for s := 0; s < num; s++ {
		name := strings.Builder{}
		name.WriteString("stu_" + strconv.Itoa(s+1)) // 根据s生成学生姓名
		stu.Name = name.String()
		scorse := randFloats(40, 100, 3) // 最低分40，最高分100，3个float64
		stu.Scores.Chinese = scorse[0]   // 语文分数
		stu.Scores.Math = scorse[1]      // 数学分数
		stu.Scores.English = scorse[2]   // 英语分数
		stu.AvgScore = stu.GenAvgScore() // 计算平均分
		students = append(students, stu)
	}
	return students
}

func main() {
	// 1.矩阵求和
	arr1 := GenArr()
	arr2 := GenArr()

	result := MatrixSum(arr1, arr2)
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			fmt.Printf("%d\t", result[i][j])
		}
		fmt.Println()
	}

	fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-")

	// 2. 根据月份判断季节
	month := 2
	Searson(month)

	fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-")

	// 3. 平均分统计
	// 3.1 统计不合格人数
	const NUMBER = 10
	students := GenStuScores(NUMBER)
	//fmt.Println(students)

	fails := map[string]Student{} // map类型存放平均分低于60的学生信息
	for _, stu := range students {
		if stu.AvgScore < 60.0 {
			fails[stu.Name] = stu
		}
	}
	fmt.Printf("平均分低于60的有%d人。\n", len(fails))

	fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-")

	// 3.2 求班级三门课平均分
	cls := Classes{Students: students}
	cls_avgs := cls.GenAvg()
	cls.Chinese_avg = cls_avgs[0]
	cls.Math_avg = cls_avgs[1]
	cls.English_avg = cls_avgs[2]

	//fmt.Println(cls)

	fmt.Printf("全班语文平均分：%.2f\n", cls_avgs[0])
	fmt.Printf("全班数学平均分：%.2f\n", cls_avgs[1])
	fmt.Printf("全班英语平均分：%.2f\n", cls_avgs[2])
}
