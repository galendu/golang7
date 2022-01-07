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
		c >>= 1
	}
	return sb.String()
}
func main() {
	//调用
	fmt.Println(BinaryFormat(123))
	fmt.Println(BinaryFormat(-123))
	fmt.Println(BinaryFormat(0))
}

// 不用分格式化输出和非格式化输出，直接使用格式化输出就行。另外只提交作业相关的内容即可
