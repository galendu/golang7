package main

import (
	"fmt"
	"os"
)

// 根据用户输入月份来判断属于那个季节
func season(m1 int) {
	switch {
	case m1 >= 1 && m1 <= 3:
		fmt.Println("当前季节是春天！")
	case m1 >= 4 && m1 <= 6:
		fmt.Println("当前季节是夏天！")
	case m1 >= 7 && m1 <= 9:
		fmt.Println("当前季节是秋天！")
	case m1 >= 10 && m1 <= 12:
		fmt.Println("当前季节是冬天！")
	default:
		fmt.Println("输入月份错误！")
		os.Exit(1)
	}

	//if m1 >=1 && m1 <=3{
	//	fmt.Println("当前季节是春天！")
	//}else if m1 >= 4 && m1 <= 6 {
	//	fmt.Println("当前季节是夏天！")
	//}else if m1 >= 7 && m1 <= 9 {
	//	fmt.Println("当前季节是秋天！")
	//}else if m1 >= 10 && m1 <= 12 {
	//	fmt.Println("当前季节是冬天！")
	//}else {
	//	fmt.Println("输入月份错误！")
	//	os.Exit(1)
	//}
}

func main() {
	// 根据输入判断季节
	var month int
	fmt.Print("请输入月份：")
	fmt.Scanln(&month)
	season(month)
}

// 完成的不错
