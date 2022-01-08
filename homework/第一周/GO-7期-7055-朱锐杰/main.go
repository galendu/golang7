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
		c >>= 1
	}
	return sb.String()
}

func main() {
	fmt.Printf("输出为%s\n", BinaryFormat(0))
	fmt.Printf("输出为%s\n", BinaryFormat(-1))
	fmt.Printf("输出为%s\n", BinaryFormat(1))
	fmt.Printf("输出为%s\n", BinaryFormat(260))
	fmt.Printf("输出为%s\n", BinaryFormat(-260))

}

// 既然格式化输出了，要让输出的信息更详细些
