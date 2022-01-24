package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

/*
	1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
	2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
*/

func arr2string(arr []int) string {
	str := strings.Builder{}
	for i := 0; i < len(arr); i++ {
		str.WriteString(strconv.Itoa(arr[i]))
		if i != len(arr)-1 {
			str.WriteString(" ")
		}
	}
	result := str.String()
	return result
}

func main() {
	// 1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
	var arr = make([]int, 0, 10)
	var map1 = make(map[int]int)

	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(128))
	}

	for _, v := range arr {
		map1[v]++
	}

	fmt.Printf("该切片中有 %d 个元素\n", len(map1))

	// 2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
	var arr1 = []int{2, 3, 4, 5}
	str := arr2string(arr1)
	fmt.Println(str)
}
