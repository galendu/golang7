package main

import (
	"fmt"
	utils "utils/binaryFormat"
)

var num int32

func binary_compute() {
	fmt.Println("=========Welcome to use binary computer============")
	fmt.Printf("0  binary: %s\n", utils.BinaryFormat(0))
	fmt.Printf("1  binary: %s\n", utils.BinaryFormat(1))
	fmt.Printf("-1 binary: %s\n", utils.BinaryFormat(-1))
	fmt.Printf("260  binary: %s\n", utils.BinaryFormat(260))
	fmt.Printf("-260 binary: %s\n", utils.BinaryFormat(-260))

	fmt.Println()
	fmt.Print("please input a number: ")
	fmt.Scanf("%d", &num)
	fmt.Printf("The binary: %v\n", utils.BinaryFormat(num))
}

func main() {
	binary_compute()
}

// 只提交作业内容接口
// 工程化实现不正确，可多学习相关内容