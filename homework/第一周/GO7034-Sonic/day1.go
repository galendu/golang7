package main

import (
	"fmt"
	"strings"
)

func BinaryFormat(n int32) string {
	a := uint32(n)
	c := uint32(1 << 31)
	sb := strings.Builder{}
	for i := 0; i < 32; i++ {
		if c&a != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1
	}
	return sb.String()
}

func main() {
	fmt.Println(BinaryFormat(0))
	fmt.Println(BinaryFormat(1))
	fmt.Println(BinaryFormat(-1))
	fmt.Println(BinaryFormat(260))
	fmt.Println(BinaryFormat(-260))
}
