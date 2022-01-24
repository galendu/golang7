package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//1.创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
func slice_Sum() {
	arr := make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(128))
	}
	// 使用map统计该切片里有多少个互不相同的元素
	var mapSlice = make(map[int]int, 100)
	for _, ele := range arr {
		v, ok := mapSlice[ele]
		if ok {
			mapSlice[ele] = v + 1
		} else {
			mapSlice[ele] = 1
		}
	}
	fmt.Printf("切片里有%d个互不相同的元素\n", len(mapSlice))
}

// 2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
func arr2String(arr []int) string {
	sb := strings.Builder{}
	for i := 0; i < (len(arr)); i++ {
		v := arr[i]
		if i == len(arr)-1 {
			sb.WriteString(fmt.Sprint(strconv.Itoa(v)))
		} else {
			sb.WriteString(fmt.Sprint(strconv.Itoa(v)))
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func main() {
	slice_Sum()
	arr := []int{2, 4, 6, 8}
	str := arr2String(arr)
	fmt.Println(str)
}

// 完成的不错，第一题想到了判断key是否存在，第二题想到了处理最后一个元素
