package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func work1() {
	rand.Seed(680)
	var tmp = make([]int, 0, 10)
	for i := 1; i <= 100; i++ {
		tmp = append(tmp, rand.Intn(128))
	}

	var countMap = make(map[int]int, 100)
	for _, ele := range tmp {
		countMap[ele] = 1
	}
	fmt.Printf("没有重复的数量有数字%d个\n", 100-len(countMap))
}

func arr2string(arr []int) string {
	sb := strings.Builder{}
	for index, ele := range arr {
		sb.WriteString(strconv.Itoa(ele))
		if index+1 != len(arr) {
			sb.WriteString(" ")
		}

	}
	merged := sb.String()
	return merged
}

func main() {
	//Q1
	work1()

	//Q2
	fmt.Println(arr2string([]int{1, 2, 3}))
}
