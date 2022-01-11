package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

/*
  1. 创建一个初始长度为0、容量为10的int型切片，
	调用rand.Intn(128)100次，往切片里面添加100个元素，
	利用map统计该切片里有多少个互不相同的元素。
*/

func total()  {
	// 定义一个len=0，cap=10的int类型的slice
	intSlice := make([]int,0,10)
	// 调用rand.Intn()函数一百次
	for i := 0; i < 100; i++ {
		intSlice = append(intSlice,rand.Intn(128))
	}
	//fmt.Println(int_Slice)

	// 定义一个map[int]bool
	totalMap := make(map[int]bool)
	// 利用map的key的唯一性来排除slice中相同元素
	for _,v := range intSlice {
		totalMap[v] = false
	}
	//fmt.Println(totalMap)

	sum:=len(totalMap)
	fmt.Printf("该切片里有%d个互不相同的元素\n",sum)

}


/*
	2. 实现一个函数func arr2string(arr []int) string，
	比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
 */

func arr2string(arr []int) string {
	sb := strings.Builder{}

	for i := 0; i < len(arr); i++ {
		str := strconv.Itoa(arr[i])
		sb.WriteString(str)
		// 判断是否是最后一个字符。如果是，则字符串末尾不加" "
		if i != len(arr) - 1 {
			sb.WriteString(" ")
		}
	}
	// TODO: 当循环完切片中最后一个值，字符串的末尾仍会加一个" "
	//for _, v := range arr {
	//	str := strconv.Itoa(v)
	//	sb.WriteString(str)
	//	sb.WriteString(" ")
	//}
	merged := sb.String()
	return merged
}

func main()  {
	total()

	var arr = []int{3,4,6}
	str := arr2string(arr)
	//fmt.Println(len(str))
	fmt.Println(str)
}