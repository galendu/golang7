package main

import (
	"fmt"
	"math"
	"strings"
)

func BinaryFormat(n int32) string {
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31))
	for i := 0; i < 32; i++ {
		if a&c > 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c = c >> 1
	}
	return sb.String()
}

func main() {
	result1 := BinaryFormat(0)
	result2 := BinaryFormat(1)
	result3 := BinaryFormat(-1)
	result4 := BinaryFormat(260)
	result5 := BinaryFormat(-260)
	fmt.Println(result1, result2, result3, result4, result5)
}

// 这种输出建议不要在一行，也考虑下格式化输出
