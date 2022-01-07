package main

import (
	"fmt"
	"math"
	"strings"
)

/*
	这个方法是32位的填充方法
*/
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
	fmt.Printf("22222的填充32位为：%s\n", BinaryFormat(222222))
	fmt.Printf("-22222的填充32位为：%s\n", BinaryFormat(-222222))
}

// 逻辑没问题，也考虑到了格式化输出，不过只提交作业要求的即可
