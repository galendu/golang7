package main

import (
	"fmt"
	"homework_1_15/homework1"
	"homework_1_15/homework2"
	"homework_1_15/homework3"
)

func main() {
	/*报错：Unable to process `evaluate`: debuggee is running
	var (
		row, column int
	)
	fmt.Print("输入数组行数和列数，使用空格分隔")
	fmt.Scanln(&row,%column)
	homework1.ArrayPlus(row, column)*/
	fmt.Println("作业一:")
	homework1.ArrayPlus(5, 8)
	fmt.Println("======================================")
	fmt.Println("作业二:")
	homework2.Season(3)
	fmt.Println("======================================")
	fmt.Println("作业三:")
	homework3.InputStudent("李四", 99.9, 77.7, 38.1)
	homework3.InputStudent("李五", 97.5, 67.3, 48.2)
	homework3.InputStudent("李六", 79.3, 87.4, 58.3)
	homework3.InputStudent("李七", 89.1, 97.6, 68.4)
	homework3.InputStudent("李八", 49.1, 37.6, 61.4)
	homework3.InputStudent("张三", 60.1, 66.6, 59.4)
	homework3.InputStudent("张四", 81.1, 55.6, 66.4)
	fmt.Println()
	fmt.Println("统计每个学生平均分及不及格人数")
	homework3.PrintOneAverage()
	fmt.Println()
	fmt.Println("统计三科平均分及每科不及格人数")
	homework3.PrintSubjectAverage("chinese")
	homework3.PrintSubjectAverage("math")
	homework3.PrintSubjectAverage("english")
	homework3.PrintSubjectAverage("oringe")
}
