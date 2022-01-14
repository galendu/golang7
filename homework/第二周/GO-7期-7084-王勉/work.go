package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func combine_str(arr []int) string {
	var str_combine = strings.Builder{}
	for _, str := range arr {
		var str_change string = strconv.Itoa(str)
		str_change += " "
		str_combine.WriteString(str_change)
	}

	fmt.Println("===================")

	return str_combine.String()

}

func main() {
	// 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
	//  实现一个函数func combine_str(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。

	var arr []int
	var sole []int
	var m map[int]int
	m = make(map[int]int)
	arr = []int{}
	sole = []int{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(128))
	}

	for _, value := range arr {
		m[value] = value
	}

	//遍历不同数
	for _, val := range m {
		sole = append(sole, val)
	}

	fmt.Println(arr)
	fmt.Println("==============================================")
	fmt.Println(m)
	fmt.Println("==============================================")
	fmt.Printf("空切片输出: %s\n", combine_str([]int{}))
	fmt.Printf("sole输出：%s\n", combine_str(sole))
	fmt.Printf("100个随机数中sole数为：%d\n", len(m)) //为何100个128随机数sole数总是69

}

// 仅输出最终结果即可
// 第一题建议使用判断map中key是否存在
// 第二题考虑处理下输出结果末尾的空格
