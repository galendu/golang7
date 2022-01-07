package main

import (
	"fmt"
	"strings"
)

//输出一个int32对应的二进制表示
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

//完成两行填空，然后写一个main函数调用BinaryFormat，看一下0、1、-1、260、-260对应的二进程表示是什么

// 可考虑格式化输出，让人一看就明白输出的是什么