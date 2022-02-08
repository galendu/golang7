package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//1.创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
func len_duplicates() {
	a := make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		a = append(a, rand.Intn(128))
	}
	fmt.Printf("a=%v\n", a)
	fmt.Printf("总数:%d\n", len(a))

	m := make(map[int]int, 100)
	for _, v := range a {
		m[v] = v
	}
	fmt.Printf("不同数字个数:%v\n", len(m))
}

//2.实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
func arr2string(arr []int) string {
	brr := make([]string, len(arr))
	for i, v := range arr {
		brr[i] = strconv.Itoa(v)
	}
	merged := strings.Join(brr, " ")
	return merged
}

func main() {
	len_duplicates()
	fmt.Println(arr2string([]int{}))
}
