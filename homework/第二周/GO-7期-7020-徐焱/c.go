package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func test() {
	sli := make([]int, 0, 10)   
	for i := 1; i <= 100; i++ { 
		rand.Seed(time.Now().UnixNano())  
		sli = append(sli, rand.Intn(128)) 
	}
	fmt.Printf("有%d个元素\n%v\n", len(sli), sli)

	var m = make(map[int]int) 
	for k, y := range sli {   
		m[y] = k 
	}

	fmt.Printf("有%d个互不相同的元素\n%v\n", len(m), m)
}

func arr2string(arr []int) string {
	sb := strings.Builder{}
	for k, v := range arr {
		if (len(arr) - 1) == k {
			sb.WriteString(strconv.Itoa(v))
		} else {
			sb.WriteString(strconv.Itoa(v))
			sb.WriteString(" ")
		}
	}
	merged := sb.String()
	return merged
}

func main() {
	test()

	arr := []int{2, 4, 6, 8, 10}
	fmt.Printf("\"%v\"\n", arr2string(arr))
}