package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//创建一个初始长度为0、容量为10的int型切片
//调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素
func slice() {
	//生成切片
	rand.Seed(time.Now().UnixNano())
	sub_slice := make([]int, 0, 10)
	sub_map := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		sub_slice = append(sub_slice, rand.Intn(128))
	}
	for _, i := range sub_slice {
		if _, exists := sub_map[i]; exists {
			sub_map[i] += 1
		} else {
			sub_map[i] = 1
		}
	}
	fmt.Printf("%d\n共有%d个不同的元素\n", sub_map, len(sub_map))
}

//实现一个函数func arr2string(arr []int) string，
//比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。

func arr2string(arr []int) string {
	sb := strings.Builder{}
	Len_arr := 0
	for _, value := range arr {
		Len_arr += 1
		sub_arr := strconv.Itoa(value)
		sb.WriteString(sub_arr)
		if Len_arr < len(arr) {
			sb.WriteString(" ")
		}

	}

	return sb.String()
}

func main() {
	slice()
	arry := []int{2, 4, 6, 8, 10, 12}
	fmt.Printf("%s", arr2string(arry))
}

// 第一个可以简化下输出
// 第二题输出可以考虑下去掉末尾的%
