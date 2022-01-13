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
