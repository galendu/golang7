package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，
//利用map统计该切片里有多少个互不相同的元素。
//2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。

func F1()  {
	var s1 = make([]int, 0, 10)
	var m1 = make(map[int]int)  // 定义一个空的map
	for i := 0;i<100;i++{
		ret := rand.Intn(128)
		s1 = append(s1, ret)
		fmt.Printf("%p  %d  %d\n", s1, len(s1), cap(s1))
		m1[ret]++   // key为添加到s1的数字，值为统计该数字出现的次数
	}
	fmt.Println("m1->", m1)
	// 校验是否准确
	//num := 0
	//for _,v:=range m1{
	//	num += v
	//}
	//fmt.Println("num=> ", num)
}

func F2(arr []int) string{
	fmt.Println("arr=>", arr)
	str := ""
	for _,v := range arr {
		str = str + strconv.Itoa(v) + " "
	}
	str = str[:len(str)-1]
	fmt.Printf("str=> %v", str)
	return str
}

func main() {
	F1()
	F2([]int{1})
}
