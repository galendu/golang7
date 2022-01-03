//这里是关于包的简单说明
package main

import "fmt"

//Add 2个整数相加
//返回和。
//
//NOTE: 注释可以有多行，但中间不能出现空行（仅有//不算空行）。
func Add(a, b int) int {
	return a + b
}

/*
Sub 函数使用示例：
  for i:=0;i<3;i++{
	  Sub(i+1, i)
  }
看到了吗？只需要行前缩进，注释里就可以写go代码，是不是很简单。
*/
func Sub(a, b int) int {
	return a - b
}

//TODO: Prod 该函数不能并发调用，需要优化
func Prod(a, b int) int {
	return a * b
}

//Deprecated: Div 不要再调用了
func Div(a, b int) int {
	return a / b
}

func 骚操作() {
	汉字变量 := "这里是值"
	fmt.Println(汉字变量)
	return
}

//main 程序入口
func main() {
	Add(3, 5)
	Sub(3, 5)
	Prod(3, 5)
	Div(3, 5)
	骚操作()
}
