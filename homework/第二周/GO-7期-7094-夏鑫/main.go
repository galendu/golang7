package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func count_slice() {
	s := make([]int, 0, 10)
	sm := make(map[int]bool)

	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(128))
	}

	for _, ele := range s {
		sm[ele] = true
	}

	fmt.Printf("该切片中有%d个不同元素\n", len(sm))
}

func arr2string(arr []int) string {
	str_build := strings.Builder{}

	num := 0

	for _, ele := range arr {
		num++
		str := strconv.Itoa(ele) //
		str_build.WriteString(str)
		if num == len(arr) {
			str_build.WriteString("")
		} else {
			str_build.WriteString(" ")
		}

	}
	return str_build.String()
}

func main() {
	/*
		1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
		2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
	*/

	// 1
	count_slice()

	// 2
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr2string(arr))
}
