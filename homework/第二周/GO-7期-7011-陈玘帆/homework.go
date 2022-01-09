package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func getCount(s []int) (int, int) {
	var m map[int]int
	m = make(map[int]int, 100)

	for _, ele := range s {
		if _, exists := m[ele]; exists{
			m[ele] += 1
		}else {
			m[ele] = 1
		}
	}

	s1 := make([]int, 0, 100)
	for key, value := range m{
		if value == 1{
			s1 = append(s1, key)
		}
	}

	return len(m), len(s1)
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
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		randint := rand.Intn(128)
		s = append(s, randint)
	}

	sc, sc2 := getCount(s)
	fmt.Printf("切片去重后的长度: %d 切片中不重复的元素个数%d\n", sc, sc2)

	fmt.Print("---将切片元素转为字符串---\n")

	s1 := make([]int, 0, 10)
	stringArr := arr2string(s1)
	fmt.Printf("原切片：%v \n转化为字符串后：%s\n", s1, stringArr)

	stringArr2 := arr2string(s)
	fmt.Printf("原切片：%v \n转化为字符串后：%s\n", s, stringArr2)

}
