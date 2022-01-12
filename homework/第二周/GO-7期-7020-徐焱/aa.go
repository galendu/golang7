package main

import (
	"fmt"
	"strconv"
	"strings"
)

func zy2(arr []int) string {
	sb := strings.Builder{}
	for k, v := range arr {
		if (len(arr) - 1) == k { //当取到最后一个index，不添加空格
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
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v", zy2(arr))
}