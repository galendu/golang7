package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

/*
创建一个初始长度为0、容量为10的int型切片，
调用rand.Intn(128)100次，往切片里面添加100个元素，
利用map统计该切片里有多少个互不相同的元素。
*/
func Homework01() {
	var sa = make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		sa = append(sa, rand.Intn(128))
	}
	// fmt.Println(len(sa))
	// fmt.Println(sa)
	var ma = make(map[int]int, 100)
	for i, v := range sa {
		ma[v] = i
	}
	fmt.Println("ma", len(ma))
}

/*
	实现一个函数func arr2string(arr []int) string，
	比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
*/

func arr2string(arr []int) string {
	sb := strings.Builder{}
	for i := 0; i < len(arr); i++ {
		value := arr[i]
		if i < len(arr)-1 { //这样避免最后多出一个空格
			sb.WriteString(fmt.Sprint(strconv.Itoa(value))) //将int转成string类型
			sb.WriteString(" ")                             //元素之间添加空格
		} else {
			sb.WriteString(fmt.Sprint(strconv.Itoa(value))) //最后一个元素
		}
	}
	sb1 := sb.String()
	return sb1
}

func main() {
	Homework01()
	arr1 := []int{}
	arr2 := []int{2, 4, 6}
	arr3 := []int{2, 4}

	fmt.Printf("arr2string(arr1)==%v\n", arr2string(arr1))
	fmt.Printf("arr2string(arr2)==%v\n", arr2string(arr2))
	fmt.Printf("arr2string(arr3)==%v\n", arr2string(arr3))
	fmt.Printf("arr2string(nil)==%v\n", arr2string(nil))
}

// 第一题建议判断map中的key是否存在，并格式化输出
// 第二题输出可以优化下
