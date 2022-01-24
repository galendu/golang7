package main

import (
	"1_8_homework/homework1"
	"1_8_homework/homework2"
	"fmt"
)

func main() {
	//作业一
	fmt.Println("=================作业一====================")
	slice := homework1.CreateSlice(100)
	homework1.PrintSlice(slice)
	fmt.Println("=================分界线====================")
	homework1.PrintMapAndLength(homework1.CountDiffKey(slice))

	//作业二
	fmt.Println("=================作业二====================")
	fmt.Println(homework2.Arr2string([]int{3, 3, 4}))
}

// 工程化实现不正确
// 第一题不用创建slice的方法就行
// 第二题有很多错误，可以修改下。也处理下结果的末尾的空格
