package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func element_count() {
	s := make([]int, 100)        //声明切片 s
	rand.Seed(time.Now().Unix()) //设置seed
	for i := 0; i < 100; i++ {
		s[i] = rand.Intn(128) //生成[0,128)区间内的随机数并存放到切片s中
	}
	m := make(map[int]string, 100) //声明哈希表 m
	for _, ele := range s {        //遍历切片s并将其element作为m的key存放，相同key会覆盖
		m[ele] = "1"
	}
	fmt.Printf("该切片有%d个不相同的元素。\n", len(m)) //统计m的key的数量
}

func arr2string(arr []int) string {
	sb := strings.Builder{} //声明一个存字符串的结构体
	for index, ele := range arr {
		s := strconv.FormatInt(int64(ele), 10) //把arr的元素转换成字符
		sb.WriteString(s)                      //追加元素字符到sb结构体
		if index == len(arr)-1 {
			break
		}
		sb.WriteString(" ")
	}
	return sb.String() //返回结构体的值
}

func call(n int) { //函数创建一个切片并调用arr2string()进行转换并输出
	s := make([]int, n)
	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}
	fmt.Printf("切片长度是%d，转换成字符串是\"%s\"\n", n, arr2string(s))
}

func main() {
	fmt.Printf("第1题结果输出：\n")
	element_count()
	fmt.Println()
	fmt.Printf("第2题结果输出：\n")
	call(5)
	call(1000)
}
