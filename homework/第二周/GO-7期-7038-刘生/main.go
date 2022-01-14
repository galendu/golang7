package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func f1() {
	a := make([]int, 0, 10)
	m := make(map[int]bool)
	count := 100
	for count > 0 {
		key := rand.Intn(128)
		a = append(a, key)
		m[key] = true
		count--
	}
	fmt.Println(len(m))
}
func arr2string(arr []int) string {
	sb := strings.Builder{}
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			sb.WriteString(strconv.Itoa(arr[i]))
		} else {
			sb.WriteString(strconv.Itoa(arr[i]))
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func main() {
	f1()
	fmt.Printf("%#v", arr2string([]int{2, 4, 6}))
}

// 第一题使用判断map的key是否存在的方式
// 第二题可以考虑去掉头尾的“以及输出的%
