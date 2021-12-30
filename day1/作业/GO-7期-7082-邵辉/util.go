package main

import (
	"fmt"
	"math"
	"strings"
)

//函数BinaryFormat 用于接收int32数据，然后以字符串形式返回二进制格式数据
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
		c >>= 1 //"1"往右移一位
	}
	return sb.String()
}

func main() {
	var a int32 = 261
	fmt.Printf("261 %s\n", BinaryFormat(a))
	fmt.Printf("-261 %s\n", BinaryFormat(-a))
}
