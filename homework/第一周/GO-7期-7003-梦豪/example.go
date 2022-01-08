package main

import (
	"fmt"
	"math"
	"strings"
)

func BinFormat(n int32) string {
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31))
	for i := 0; i < 32; i++ {
		if a&c != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		// fmt.Println(c)
		c >>= 1
	}
	return sb.String()
}

func main() {
	// a := BinFormat(10000)
	// fmt.Println(a)
	fmt.Println(BinFormat(0))
	fmt.Println(BinFormat(1))
	fmt.Println(BinFormat(-1))
	fmt.Println(BinFormat(260))
	fmt.Println(BinFormat(-260))

}

// 结果是正确的，可以尝试下格式化输出