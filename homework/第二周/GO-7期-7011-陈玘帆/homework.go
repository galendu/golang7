package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func getSliceDistinctElementsCount(s []int) int {
	var m map[int]bool
	m = make(map[int]bool, 100)

	for _, ele := range s {
		m[ele] = true
	}

	return len(m)
}


func arr2string(arr []int) string{
	sb := strings.Builder{}

	for _, ele := range arr{
		var s string = strconv.Itoa(ele)
		sb.WriteString(s)
		sb.WriteString(" ")
	}
	mergedSb := sb.String()
	return mergedSb
}


func main() {
	fmt.Print("---获取切片中互不相同的元素个数---\n")
	s := make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		randint := rand.Intn(128)
		s = append(s, randint)
	}

	sl := getSliceDistinctElementsCount(s)
	fmt.Printf(" slice distinct elements count: %d\n", sl)

	fmt.Print("---将切片元素转为字符串---\n")

	s1 := make([]int, 0, 10)
	stringArr := arr2string(s1)
	fmt.Printf("原切片：%v \n转化为字符串后：%s\n", s1, stringArr)

	stringArr2 := arr2string(s)
	fmt.Printf("原切片：%v \n转化为字符串后：%s\n", s, stringArr2)

}
