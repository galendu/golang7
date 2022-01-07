package main

import (
	"fmt"
	"math"
	"strings"
)

//输出一个int32对应的二进制表示
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
		c = c >> 1
	}
	return sb.String()
}

func main() {
	fmt.Println("hello World")
	fmt.Printf("2的二进制：%s\n", BinaryFormat(2))
	fmt.Printf("-2的二进制：%s\n", BinaryFormat(-2))
	fmt.Printf("10的二进制：%s\n", BinaryFormat(10))
	fmt.Printf("321的二进制：%s\n", BinaryFormat(321))
	fmt.Printf("1987的二进制：%s\n", BinaryFormat(1987))
	fmt.Printf("-1000的二进制：%s\n", BinaryFormat(-1000))
	fmt.Printf("-32的二进制：%s\n", BinaryFormat(-32))
	fmt.Printf("9999的二进制：%s\n", BinaryFormat(9999))
}
