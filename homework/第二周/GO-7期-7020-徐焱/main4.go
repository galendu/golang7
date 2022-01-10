package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func zy1() {
	sli := make([]int, 0, 10)   //定义一个长度为0、容量为10的int型切片
	for i := 1; i <= 100; i++ { //循环100次
		rand.Seed(time.Now().UnixNano())  //设置随机数
		sli = append(sli, rand.Intn(128)) //将随机数append到sli切片
	}
	fmt.Printf("有%d个元素\n%v\n", len(sli), sli)

	var m = make(map[int]int) //定义一个map
	for k, y := range sli {   //遍历sli
		m[y] = k //将sli的值指定为map的index，因为index不能重复，若重复则追加
	}

	fmt.Printf("map中有%d个互不相同的元素\n%v\n", len(m), m)
}

func zy2(arr []int) string {
	sb := strings.Builder{}
	for k, v := range arr {
		if (len(arr) - 1) == k { //当取到最后一个index，不添加空格
			sb.WriteString(strconv.Itoa(v))
		} else {
			sb.WriteString(strconv.Itoa(v))
			sb.WriteString(" ")

		}
	}
	merged := sb.String()
	return merged
}

func main() {
	//1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次
	zy1()

	//2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}
	arr := []int{2, 4, 6, 8, 10}
	fmt.Printf("返回\"%v\"\n", zy2(arr))
}