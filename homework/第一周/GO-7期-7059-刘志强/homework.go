package main

import (
	"fmt"
	"math"
	"strings"
)

func BinaryFormat(n int32) string {
	a := uint32(n)
	fmt.Println(a)
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

//完成两行填空，然后写一个main函数调用BinaryFormat，看一下0、1、-1、260、-260对应的二进程表示是什么
func main() {
	fmt.Println(BinaryFormat(0))
	fmt.Println(BinaryFormat(1))
	fmt.Println(BinaryFormat(-1))
	fmt.Println(BinaryFormat(260))
	fmt.Println(BinaryFormat(-260))
}

// 1、可以格式化输出，让人很容易看明白
// 2、提交作业时，可以不用提交中间打印验证过程
