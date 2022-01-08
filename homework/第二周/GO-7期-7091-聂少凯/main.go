package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func arr2string(arr []int) string {
	var (
		tmpStr string
		sb     strings.Builder
	)
	for _, i := range arr {
		tmpStr = strconv.Itoa(i)
		sb.WriteString(tmpStr + " ")
	}
	return sb.String()
}

func arr2stringTest() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr2string(arr))
}

func MapVertifyElements() {
	Slice := make([]int, 0, 10)
	Map := make(map[int]int, 10)
	for i := 0; i <= 100; i++ {
		tmp := rand.Intn(128)
		Slice = append(Slice, tmp)
		Map[tmp]++
	}
	fmt.Printf("我发现map里有%d个互不相同的元素\n", len(Map))
}

func main() {
	//实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
	arr2stringTest()

	//创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
	MapVertifyElements()
}
