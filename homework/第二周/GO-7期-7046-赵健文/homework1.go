// 1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genSlice(c int) []int {
	s := make([]int, 0, 10)      // 初始化切片
	rand.Seed(time.Now().Unix()) // 设定以unix当前时间戳随机种子，每次调用都产生不同随机数
	for i := 0; i < c; i++ {
		s = append(s, rand.Intn(128))
	}
	return s
}
func uniqCount(s []int) int {
	m := make(map[int]bool, len(s)) // 初始化map，并指定长度，避免系统频繁申请内存
	for _, ele := range s {
		m[ele] = true // 利用map的key去重，value使用bool值，bool占用1个字节，节省内存
	}
	return len(m)
}

func main() {
	s := genSlice(100)
	c := uniqCount(s)
	fmt.Printf("There are %d different elements\n", c)
}

