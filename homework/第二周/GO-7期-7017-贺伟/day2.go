package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func slice_int() {
	s := make([]int, 0, 10)      //声明初始长度为0，容量为10的int切片
	rand.Seed(time.Now().Unix()) //使每次得到的随机数不一样
	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(128))
	}
	fmt.Printf("s=%v\n", s)     //打印切片
	m := make(map[int]int, 100) //声明map
	for _, i := range s {
		if _, ok := m[i]; !ok {
			m[i] = 1
		} else {
			m[i]++
		}
	}
	num := 0
	for key, value := range m {
		if m[key] == 1 {
			num += 1
		}
		fmt.Println(key, value)
	}
	fmt.Printf("一共有%d个不同元素\n", num)
}

//
func arr2string(arr []int) string {
	sb := strings.Builder{}
	for i := 0; i < len(arr); i++ {
		sb.WriteString(strconv.Itoa(arr[i])) //将取到的值转换成string后加入到sb中
		sb.WriteString(" ")
	}
	return sb.String()
}

func main() {
	slice_int()
	a := []int{1, 2, 3, 4, 10, 20, 30}
	fmt.Printf("拼接a=%s\n", arr2string(a))
	fmt.Println("1")
}
