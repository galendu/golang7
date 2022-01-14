package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func test1() {
	s := make([]int, 0, 10)
	m := make(map[int]int, 50)
	for i := 0; i < 100; i++ {
		r := rand.Intn(128)
		s = append(s, r)
		if v, exist := m[r]; exist {
			m[r] = v + 1
		} else {
			m[r] = 1
		}
	}

	for i, v := range s {
		fmt.Printf("%d -> %d \n", i, v)
	}

	for k, v := range m {
		fmt.Printf("%d -> %d\n", k, v)
	}
}

func arr2string(arr []int) string {
	strBuilder := strings.Builder{}
	for i, e := range arr {
		strBuilder.WriteString(strconv.Itoa(e))
		if i < len(arr)-1 {
			strBuilder.WriteString(" ")
		}
	}
	return strBuilder.String()
}

func main() {
	test1()

	s := []int{1, 2, 3, 4, 5, 6, 7}
	result := arr2string(s)
	fmt.Println(result)
}
