package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetMatrixSum(){
	const Row, Col = 8, 5
	m1 := [Row][Col]int{}
	m2 := [Row][Col]int{}
	m3 := [Row][Col]int{}

	rand.Seed(time.Now().Unix())
	for i:=0;i<Row;i++{
		for j:=0;j<Col;j++{
			m1[i][j] = rand.Intn(128)
			m2[i][j] = rand.Intn(128)
		}
	}

	fmt.Println(m1)
	fmt.Println(m2)

	for i:=0;i<Row;i++{
		for j:=0;j<Col;j++{
			m3[i][j] = m1[i][j] + m2[i][j]
		}
	}

	fmt.Println(m3)
}


func GetSeasonByIf(seasonInt int){
	if seasonInt >= 3 && seasonInt <= 5{
		fmt.Println("春季")
	}else if seasonInt >= 6 && seasonInt <= 8{
		fmt.Println("夏季")
	}else if seasonInt >= 9 && seasonInt <= 11 {
		fmt.Println("秋季")
	}else if seasonInt == 12 || seasonInt == 1 || seasonInt == 2{
		fmt.Println("冬季")
	}else {
		fmt.Println("错误的月份")
	}
}

func GetSeasonBySwitch(seasonInt int){
	switch {
	case seasonInt >= 3 && seasonInt <= 5:
		fmt.Println("春季")
	case seasonInt >= 6 && seasonInt <= 8:
		fmt.Println("夏季")
	case seasonInt >= 9 && seasonInt <= 11:
		fmt.Println("秋季")
	case seasonInt == 12 || (seasonInt >= 1 && seasonInt <= 2):
		fmt.Println("冬季")
	default:
		fmt.Println("错误的月份")
	}
}

type student struct {
	name string
	chinese int
	math int
	english int
}

func (stu student) studentGpa() float64{
	gpa := float64(stu.chinese + stu.math + stu.english) / 3
	return gpa
}

type class struct {
	chinese_gpa float64
	math_gpa float64
	english_gpa float64
	students []student
}

func (cls *class) setClassGpa(course string){
	totalScore := 0
	studentCount := len(cls.students)
	for _, stu := range cls.students{
		switch course {
		case "chinese":
			totalScore += stu.chinese
		case "math":
			totalScore += stu.math
		case "english":
			totalScore += stu.english
		}
	}
	gpa := float64(totalScore / studentCount)

	switch course {
	case "chinese":
		cls.chinese_gpa = gpa
	case "math":
		cls.math_gpa = gpa
	case "english":
		cls.english_gpa = gpa
	}
}

func (cls class) gpaLower60StudentsCount() int{
	count := 0
	for _, stu := range cls.students{
		stu_gpa := stu.studentGpa()
		if stu_gpa < 60{
			count += 1
		}
	}
	return count
}

func generateClass() class{
	cls := class{}

	// 给班级填充50个学员
	rand.Seed(time.Now().Unix())
	for i:=0; i< 50; i++ {
		stu := student{name: "学生" + strconv.Itoa(i), math: rand.Intn(101), chinese: rand.Intn(101),
			english: rand.Intn(101)}
		cls.students = append(cls.students, stu)
		fmt.Printf("学员：%s，语文:%d，数学:%d, 英语:%d，平均分:%0.2f\n", stu.name, stu.chinese, stu.math,
			stu.english, stu.studentGpa())
	}

	cls.setClassGpa("chinese")
	cls.setClassGpa("math")
	cls.setClassGpa("english")
	return cls
}

func showClassInfo(){
	cls := generateClass()
	student_count := cls.gpaLower60StudentsCount()
	fmt.Printf("--班级语文平均分:%0.2f，英语平均分:%0.2f，数学平均分:%0.2f--\n", cls.chinese_gpa, cls.english_gpa, cls.math_gpa)
	fmt.Printf("--班级中平均分低于60的学生个数有%d个--\n", student_count)

}


func main(){
	fmt.Print("---随机初始化两个8*5的矩阵，求两个矩阵的和---\n")
	GetMatrixSum()

	fmt.Print("---给定月份，判断属于哪个季节。分别用if和switch实现---\n")
	fmt.Print("---if---\n")
	GetSeasonByIf(2)
	GetSeasonByIf(6)
	GetSeasonByIf(3)
	GetSeasonByIf(10)
	GetSeasonByIf(100)
	fmt.Print("---switch---\n")
	GetSeasonBySwitch(2)
	GetSeasonBySwitch(6)
	GetSeasonBySwitch(3)
	GetSeasonBySwitch(10)
	GetSeasonBySwitch(100)

	fmt.Print("---创建一个Student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，" +
		"求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位---\n")
	showClassInfo()
}
