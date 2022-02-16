package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func one() {
	s := make([]int, 0, 10)
	m := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(128))
	}
	fmt.Printf("len %d\ns=%v\n", len(s), s)
	for i := 0; i < len(s); i++ {
		key := s[i]
		if _, exits := m[key]; exits {
			m[key] += 1
		} else {
			m[key] = 1
		}
	}
	fmt.Printf("互不相同的元素有%d个\nm=%v\n", len(m), m)
}

func arr2string(arr []int) string {
	if len(arr) == 0 {
		return ""
	} else {
		s := strings.Builder{}
		for ind, ele := range arr {
			sele := strconv.Itoa(ele)
			s.WriteString(sele)
			if ind != len(arr)-1 {
				s.WriteString(" ")
			}
		}
		return s.String()
	}
}

func main() {
	// one()
	fmt.Printf("%se\n", arr2string([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
