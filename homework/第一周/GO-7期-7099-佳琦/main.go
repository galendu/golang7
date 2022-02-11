package main

import (
	"fmt"
	"math"
	"strings"
)

// ⭐️ 输出一个int32对应的二进制表示
// 基本思路: c的二进制, 从1000..000(2^31), 到 000...001(2^0)
func BinaryFormat(n int32) string {
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31))
	// fmt.Println(a, c)
	for i := 0; i < 32; i++ {
			if  a & c > 1 {
					sb.WriteString("1")
			} else {
					sb.WriteString("0")
			}
			c >>= 1
			// fmt.Printf("c: %d\n", c)
	}
	return sb.String()
}

func BinaryFormat8(n int8) string {
	a := uint8(n)
	sb := strings.Builder{}
	c := uint8(math.Pow(2, 7))
	for i := 0; i < 8; i++ {
		if a & c > 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1
	}
	return sb.String()
}

func PrintBinStr8(n int8) {
	binStr := BinaryFormat8(n)
	fmt.Printf("%d binary string: %s\n", n, binStr)
}

func PrintBinStr(n int32) {
	binStr := BinaryFormat(n)
	fmt.Printf("%d binary string: %s\n", n, binStr)
}

// 完成两行填空，然后写一个main函数调用BinaryFormat，看一下0、1、-1、260、-260对应的二进程表示是什么

func main() {
	var i int8 = -128
	for ; i <= 127; i++ {
		PrintBinStr8(i)
		if i == 127 {
			break
		}
	}
	// PrintBinStr8(-8)
	// PrintBinStr8(-1)
	// PrintBinStr8(0)
	// PrintBinStr8(1)
	// PrintBinStr(8)
	// PrintBinStr(260)
	// PrintBinStr(-260)
}
