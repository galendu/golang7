package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func task() {
	arr := make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(128))
	}
	m := make(map[int]int, len(arr))
	for i, ele := range arr {
		m[ele] = arr[i]
	}
	fmt.Printf("len of map is,%d\n", len(m))
}

func arr2string(arr []int) string {
	sb := strings.Builder{}
	if arr != nil {
		for ele := range arr {
			sb.WriteString(strconv.Itoa(ele))
			sb.WriteString(" ")
		}
		return strings.Trim(sb.String(), " ")
	} else {
		return "999999-切片无值"
	}
}

func main() {
	task()
	arr := make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(128))
	}
	fmt.Println(arr2string(arr))
}
