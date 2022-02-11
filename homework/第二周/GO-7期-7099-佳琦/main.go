// 1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
// 2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func test1() map[int]int {
	s := make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(128))
	}
	m := make(map[int]int)
	for _, val := range(s) {
		count, _ := m[val];
		m[val] = count + 1
	}
	return m
}

func arr2string(arr []int) string {
	sb := strings.Builder{}
	for _, val := range(arr) {
		str := strconv.Itoa(val) + " "
		sb.WriteString(str)
	}
	return sb.String()
}

func main() {
	m := test1()
	fmt.Println(m)
	str := arr2string([]int{1, 2, 3})
	fmt.Println(str)
}