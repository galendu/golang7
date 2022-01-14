/*
1. 创建一个初始长度为0、容量为10的int型切片
调用rand.Intn(128)100次，往切片里面添加100个元素
利用map统计该切片里有多少个互不相同的元素。
2. 实现一个函数func arr2string(arr []int) string
比如输入[]int{2,4,6}，返回“2 4 6”
输入的切片可能很短，也可能很长。
*/

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// 实现一个函数func arr2string(arr []int) string
func arr2string(arr []int) string {

	s := strings.Builder{}
	s.WriteString("“")

	//将arr转换成字符串数组
	for _, v := range arr {
		a := strconv.Itoa(v)

		s.WriteString(a)
		s.WriteString(" ")

	}

	s.WriteString("”")
	return s.String()
}

func main() {

	//创建一个初始长度为0、容量为10的int型切片
	s := make([]int, 0, 10)

	//往切片里添加100个元素
	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(128))
	}

	//打印s
	//fmt.Println(s)

	//利用map统计该切片里有多少个互不相同的元素

	//初始化map
	m := make(map[int]int, 100)

	//统计
	for _, v := range s {
		m[v] += 1
	}
	//fmt.Println(m)
	fmt.Printf("切片s中互不相同的元素个数为：%d\n", len(m))
	//fmt.Println("每个元素重复的个数为：")
	//for v, i := range m {
	//	fmt.Printf("元素%d：%d\n", v, i)
	//}

	//调用arr2string函数，参数为：[]int{2,4,6}，返回值为：“2 4 6”

	fmt.Printf("arr2string([]int{2,4,6})=%s", arr2string([]int{2, 4, 6}))

}
