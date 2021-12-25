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
		if a&c > 0 { // 判断a当前位上是否为1
			sb.WriteString("1") // &运算对应位上都为1，结果为1
		} else {
			sb.WriteString("0")
		}
		c >>= 1 // c右移一位
	}
	return sb.String()
}

func main() {
	fmt.Printf("0: %s\n", BinaryFormat(0))

	fmt.Printf("1: %s\n", BinaryFormat(1))
	fmt.Printf("-1: %s\n", BinaryFormat(-1))

	fmt.Printf("65535: %s\n", BinaryFormat(65535))
	fmt.Printf("-65535: %s\n", BinaryFormat(-65535))

	fmt.Printf("2147483647: %s\n", BinaryFormat(2147483647))
	fmt.Printf("-2147483647: %s\n", BinaryFormat(-2147483647))
}
