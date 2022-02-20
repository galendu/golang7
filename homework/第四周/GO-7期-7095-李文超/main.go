package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	//ss()
	arr := []int{2, 2, 3, 4, 3}
	fmt.Println(arr2string(arr))

}

//1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
func ss() {
	qp := make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		qp = append(qp, rand.Intn(128))
	}
	//fmt.Printf("切片长度为:%d,容量为:%d\n",len(qp),cap(qp))
	//fmt.Printf("内容为:%d\n",qp)

	map1 := make(map[int]int, 100)
	for _, err := range qp {
		map1[err] = 1
	}
	fmt.Printf("有%d个互不相同的元素", len(map1))
}

//2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
func arr2string(arr []int) string {
	sb := strings.Builder{}
	for index := 0; index < len(arr)-1; index++ {
		sb.WriteString(strconv.Itoa(arr[index]))
		sb.WriteString(" ")
	}
	indexlast := len(arr) - 1
	//fmt.Printf("长度为%d,类型为%T",indexlast,indexlast)
	sb.WriteString(strconv.Itoa(arr[indexlast]))
	result := sb.String()
	return result
}
