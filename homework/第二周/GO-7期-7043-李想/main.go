package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//作业1函数
func staticEle()int{
	arr := make([]int,0,10)
	for i := 0;i < 100;i++{
		arr = append(arr,rand.Intn(128))
	}

	m := make(map[int]bool,100)
	for _,ele := range arr{
		m[ele] = true
	}
	count := len(m)

	return count
}

//作业2函数
func arr2string(arr []int)string{
	s := strings.Builder{}
	for _,ele := range arr{
		s.WriteString(strconv.Itoa(ele))
		s.WriteString(" ")
	}
	merged := s.String()
	return merged
}

func main(){
	//作业1运行代码
	//count := staticEle()
	//fmt.Println(count)

	//作业2运行代码
	arr := []int{1,2,3}
	merged := arr2string(arr)
	fmt.Println(merged)
}

