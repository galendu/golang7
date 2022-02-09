package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	// "math/rand"
)

func get_int() {
	cur_int := make([]int, 0, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {

		cur_int = append(cur_int, rand.Intn(128))

	}
	count_int := make(map[int]int)
	for _, value := range cur_int {
		count_int[value] = count_int[value] + 1
	}

	fmt.Printf("随机100个元素中，共有%d不同的元素", len(count_int))
}

func arr2string(arr []int) string {
	toStr := strings.Builder{}
	var b string
	for _, value := range arr {
		a := strconv.Itoa(value)
		b := strings.Join([]string{a, b}, "  ")
		toStr.WriteString(b)
	}
	return strings.TrimSuffix(toStr.String(), " ")

}

func main() {

	//统计100个元素中有多少个互不相同元素
	get_int()

	//2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
	arr := []int{2, 3, 20, 103, 10}
	fmt.Printf("%v", arr2string(arr))

}