package main

import (
	"fmt"
	"strings"
    "strconv"

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
	merged := s.String()
	return merged
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Printf("\"%v\"\n", arr2string(arr))
}