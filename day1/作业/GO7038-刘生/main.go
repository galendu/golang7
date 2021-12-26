package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(BinaryFormat(0))
	fmt.Println(BinaryFormat(1))
	fmt.Println(BinaryFormat(-1))
	fmt.Println(BinaryFormat(260))
	fmt.Println(BinaryFormat(-260))
}

//输出一个int32对应的二进制表示
func BinaryFormat(n int32) string {
	a := uint32(n)
	var c uint32
	sb := strings.Builder{}
	for i := 0; i < 32; i++ {
		c = 1 << (32 - i - 1)
		if c&a != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}
	return sb.String()
}

//完成两行填空，然后写一个main函数调用BinaryFormat，看一下0、1、-1、260、-260对应的二进程表示是什么
