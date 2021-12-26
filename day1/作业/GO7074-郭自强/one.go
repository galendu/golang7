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
		//判断当前位是否为1
		if a&c != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		//右移1位
		c >>= 1
	}
	return sb.String()
}

/*
完成两行填空，然后写一个main函数调用BinaryFormat，
看一下0、1、-1、260、-260对应的二进程表示是什么

*/

//调用BinaryFormat
func main() {
	fmt.Printf("113=%s\n", BinaryFormat(113))
	fmt.Printf("-113=%s\n", BinaryFormat(-113))
	fmt.Printf("648=%s\n", BinaryFormat(648))
	fmt.Printf("-648=%s\n", BinaryFormat(-648))
	fmt.Printf("1226=%s\n", BinaryFormat(1226))
	fmt.Printf("-1226=%s\n", BinaryFormat(-1226))

}
