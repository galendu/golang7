package main

import (
	"fmt"
	"os"
)

type student struct {
	Name        string // 姓名
	Language    int    // 语文
	Mathematics int    // 数学
	English     int    // 英语
}

// 存储学生信息
var studentSlice = make([]student, 0, 10)

// student 构造函数
func newStudent(name string, language, mathematics, english int) student {
	return student{
		Name:        name,
		Language:    language,
		Mathematics: mathematics,
		English:     english,
	}
}

// student方法 求每个人的平均分数
func (s student) studentAverage() {
	averageFail := 0
	for _, v := range studentSlice {
		s1 := (v.Language + v.Mathematics + v.English) / 3
		fmt.Printf("%s 三科平均成绩：%d\n", v.Name, s1)
		if s1 <= 60 {
			averageFail++
		}
	}
	fmt.Printf("平均分低于60的有：%d个\n", averageFail)
}

// student方法 求全班同学三科成绩平均值
func (s student) studentAllAverage() {
	sumLanguage := 0
	sumMathematics := 0
	sumEnglish := 0
	for _, i := range studentSlice {
		sumLanguage += i.Language
		sumMathematics += i.Mathematics
		sumEnglish += i.English
	}
	sumLanguageAver := sumLanguage / len(studentSlice)
	sumMathAver := sumMathematics / len(studentSlice)
	sumEngAver := sumEnglish / len(studentSlice)
	fmt.Printf("全班三科平均值,语文：%d 数学：%d 英语：%d\n", sumLanguageAver, sumMathAver, sumEngAver)
}

// student方法 查看学生信息列表
func (s student) studentList() {
	for _, v := range studentSlice {
		fmt.Printf("姓名：%s 语文成绩：%d 数学成绩：%d 英语成绩：%d\n", v.Name, v.Language, v.Mathematics, v.English)
	}
}

// student方法 新增学生信息
func (s student) studentAdd(num int) {
	var (
		//stuId       int
		name        string
		language    int
		mathematics int
		english     int
	)

	//fmt.Print("请输入学号：")
	//fmt.Scanln(&stuId)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入语文成绩：")
	fmt.Scanln(&language)
	fmt.Print("请输入数学成绩：")
	fmt.Scanln(&mathematics)
	fmt.Print("请输入英语成绩：")
	fmt.Scanln(&english)
	s1 := newStudent(name, language, mathematics, english)
	studentSlice = append(studentSlice, s1)
}

func main() {
	// 学生信息管理系统
	num := 1
	for {
		s1 := student{}
		fmt.Println("欢迎光临学生信息管理系统！")
		fmt.Println(`
		1、查看学生信息
		2、新增学生信息
		3、学生成绩平均
		4、全班成绩平均
		5、推出系统`,
		)
		fmt.Print("请输入序号：")
		var number int
		fmt.Scanln(&number)
		//fmt.Printf("输入的序号是：%d\n", number)
		switch number {
		case 1:
			s1.studentList()
		case 2:
			s1.studentAdd(num)
			num++
		case 3:
			s1.studentAverage()
		case 4:
			s1.studentAllAverage()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入错误！")
		}
	}
}

// 实现作业要求即可
