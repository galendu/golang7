package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func statistics_slice(n int) {
	rand.Seed(time.Now().UnixMilli())
	slice := make([]int, 0, 10) // 声明slice用于存储数据
	m := make(map[int]int, n)   // 声明map，用于统计各个值出现的次数
	for i := 0; i < n; i++ {
		num := rand.Intn(128)      // 获取随机数
		slice = append(slice, num) // 写入slice
		m[num] += 1                // 通过每次的随机数做key，每取到一次对应的value + 1
	}

	fmt.Printf("There are %d different elements in the slice.\n", len(m)) // map的len就是不同数值出现的次数

	// 遍历map, 展示每个值出现的次数
	/*
		for key, value := range m {
			fmt.Printf("数字%d出现了%d次\n", key, value)
		}
	*/
}

func arr2string(arr []int) string {
	if len(arr) == 0 || arr == nil {
		return "Please pass in a slice that is not empty or nil!"
	} else {
		result := strings.Builder{}

		for i := 0; i < len(arr); i++ {
			//str := strconv.Itoa(arr[i]) + " " // + " "为了输出好看，但是最后一个元素末尾会有" "
			str := strconv.FormatInt(int64(arr[i]), 10) + " " // 性能会更高
			result.WriteString(str)
		}
		return strings.TrimSpace(result.String()) // strings.TrimSpace()去掉两边的空白字符
	}
}

func main() {
	// 1.创建一个初识长度为0、容量为10的整型切片，调用rand.Intn(128) 100次，往切片里添加100个元素，利用map统计该切片有多少个互不相同的元素
	statistics_slice(100)

	// 2.实现一个函数func arr2string(arr []int)string，比如输入[]int{2, 4, 6}，返回 "2 4 6"，输入的的切片可能很短，也可能很长

	slice := []int{2, 4, 6}
	//slice := []int{2, 4, 6, 1, 7, 9, 1, 8, 12, 15, 38}
	//slice := []int{2}
	//slice := []int{}
	//slice := make([]int, 0)
	//str := arr2string(nil)
	str := arr2string(slice)
	fmt.Printf("[%s]\n", str)
}

// 第一题建议使用判断map中是否存在key的方式
// 第二题结果不是要求的，可以更改下
