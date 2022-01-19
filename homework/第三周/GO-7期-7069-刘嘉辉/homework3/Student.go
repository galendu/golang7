package homework3

import "fmt"

//自定义学生类
type Student struct {
	name                   string  //student name
	chinese, math, english float32 //score
	average                float32 //平均值
}

//切片初始化
var allStudent = []Student{}

//学生构造函数
func newStudent(name string, chineseScore, mathScore, englishScore float32) *Student {
	student := Student{
		name:    name,
		chinese: chineseScore,
		math:    mathScore,
		english: englishScore,
		//计算平均值
		average: (chineseScore + mathScore + englishScore) / 3,
	}
	allStudent = append(allStudent, student)
	return &student
}

//输入学生数据
func InputStudent(name string, chineseScore, mathScore, englishScore float32) {
	//调用构造函数
	var stu *Student = newStudent(name, chineseScore, mathScore, englishScore)
	//打印输入数据
	fmt.Printf("%s 学生语文成绩：%.2f，数学成绩：%.2f，英语成绩：%.2f\n", stu.name, stu.chinese, stu.math, stu.english)

	/*
			测试代码
			fmt.Printf("%T\n",*stu)
			fmt.Printf("%T\n",stu)
			fmt.Println(*stu)
			fmt.Println(stu)
		    var AllStudent []Student = make([]Student, 20)
			fmt.Printf("%T",AllStudent)
			AllStudent=append(AllStudent,*stu)
	*/
}

//输出某学生学生三科平均值及不及格人数
func PrintOneAverage() {
	noPassNum := 0
	for _, value := range allStudent {
		if value.average < 60 {
			noPassNum++
		}
		fmt.Printf("%s学生三科平均分是%.2f\n", value.name, value.average)
	}
	fmt.Printf("三科平均值小于60分的有%d人\n", noPassNum)
}

func PrintSubjectAverage(subject string) {
	var (
		projectAverage float32
		noPassNum      int
	)
	if subject == "chinese" {
		for _, value := range allStudent {
			projectAverage += value.chinese
			if value.chinese < 60 {
				noPassNum += 1
			}
		}
	} else if subject == "math" {
		for _, value := range allStudent {
			projectAverage += value.math
			if value.math < 60 {
				noPassNum += 1
			}
		}
	} else if subject == "english" {
		for _, value := range allStudent {
			projectAverage += value.english
			if value.english < 60 {
				noPassNum += 1
			}
		}
	} else {
		fmt.Println("检查学科输入是否正确（chinese/math/english）")
		return
	}
	fmt.Printf("%s学科班级平均分是%.2f分，该学科共有%d人不及格\n", subject, projectAverage/float32(len(allStudent)), noPassNum)
}
