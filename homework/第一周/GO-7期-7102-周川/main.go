//完成两行填空，然后写一个main函数调用BinaryFormat，看一下0、1、-1、260、-260对应的二进程表示是什么
package main

import (
	"fmt"
	"math"
	"strings"
)

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
		c = c >> 1
	}
	return sb.String()
}

func main() {
	fmt.Printf("0=%s\n1=%s\n-1=%s\n260=%s\n-260=%s\n", BinaryFormat(0), BinaryFormat(1), BinaryFormat(-1), BinaryFormat(260), BinaryFormat(-260))
}

