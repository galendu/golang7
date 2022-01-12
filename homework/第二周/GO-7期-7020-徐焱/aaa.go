package main

import (
	"fmt"
	"strconv"
	"strings"
)

func arr2string(arr []int) string {
	s := strings.Builder{}
	for k, v := range arr {
		if (len(arr) - 1) == k {
			s.WriteString(strconv.Itoa(v))
		} else {
			s.WriteString(strconv.Itoa(v))
			s.WriteString(" ")
		}
	}
	m := s.String()
	return m
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v", arr2string(arr))
}