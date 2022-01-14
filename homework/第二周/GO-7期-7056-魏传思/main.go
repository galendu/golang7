package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//作业1 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次
//往切片里添加100个元素，利用map统计该切片里有多少个互不相同的元素
func ar(n int) []int {
	a := make([]int, 100)
	for i := 0; i < 100; i++ {
		a = append(a, rand.Intn(128))
	}
	return a
}

func mapUse(ar []int) int {
	m := make(map[int]bool, 200)
	for _, v := range ar {
		m[v] = false
	}
	return len(m)
}

//作业2 实现一个函数func arr2string(arr []int) string
//比如输入[]int{2,4,6},返回"2 4 6".
//输入的切片可能很短，也可能很长
func arr1string(arr []int) string {
	var sb strings.Builder
	for _, v := range arr {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteString(" ")
	}
	return sb.String()
}

func main() {
	s := ar(100)
	m := mapUse(s)
	fmt.Println(m)

	arr := arr1string([]int{1, 2, 3, 5})
	fmt.Println(arr)
}
