package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//统计100个元素中有多少个互不相同元素
func diffEle(number int) {
	s := make([]int, 0, 10)
	rand.Seed(time.Now().UnixNano()) //解决rand.Intn函数每次都返回同样随机数的问题
	for i := 0; i < number; i++ {
		s = append(s, rand.Intn(128)) //调用rand.Intn函数从0-127的值中添加1个元素到slice s
	}
	m := make(map[int]int, number)
	for _, i := range s { //遍历slice
		if _, exists := m[i]; exists { //判断是否有相同的元素，如果有则value+1
			m[i] += 1
		} else {
			m[i] = 1
		}
	}
	var num int
	for key := range m { //遍历map
		if m[key] == 1 { //value为1的元素为互不相同，num+1，num最后的值即为总互不相同元素数
			num += 1
		}
	}
	fmt.Printf("一共有%d个互不相同的元素\n", num)
}

func arr2string(arr []int) string {
	as := strings.Builder{}
	count := 0 //计数器
	as.WriteString("\"")
	for _, value := range arr {
		count++
		str := strconv.Itoa(value)
		as.WriteString(str)
		if count != len(arr) { //如果count等于arr的len值，说明是最后一次循环，则不添加" "
			as.WriteString(" ")
		}
	}
	as.WriteString("\"")
	return as.String()
}

func main() {
	/*
		作业：
		1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
		2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
	*/

	diffEle(100)
	a := []int{1, 25, 6, 164, 33, 515, 6, 1}
	fmt.Printf("%s\n", arr2string(a))
}

// 第一题思路可以优化下
// 第二题可以想下把头尾的"去掉
