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
		if a&c != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1
	}
	return sb.String()
}

func main() {
	fmt.Printf("os arch %s\n", BinaryFormat(160))
	fmt.Printf("os arch %s\n", BinaryFormat(-160))
	fmt.Printf("os arch %s\n", BinaryFormat(16))
	fmt.Printf("os arch %s\n", BinaryFormat(-16))
}

// 逻辑正确，格式化输出时内容要对应
