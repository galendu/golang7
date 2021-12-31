package main

import (
	"fmt"
	"math"
	"strings"
)

func BinaryFormat(n int32) string {
	a:=uint32(n)
	sb:=strings.Builder{}
	c:=uint32(math.Pow(2,31))
	for i := 0; i < 32; i++ {
		if a&c !=0 {
			sb.WriteString("1")
		}else {
			sb.WriteString("0")
		}
		c >>= 1
	}
	return sb.String()
}
func main() {
	a:=BinaryFormat(32)
	b:=BinaryFormat(-128)
	c:=BinaryFormat(0)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
