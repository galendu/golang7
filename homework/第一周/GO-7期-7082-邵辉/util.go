package main

import (
	"fmt"
	"math"
	"strings"
)

//函数BinaryFormat 将int32转换为二进制
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
	fmt.Printf("   0的二进制是 %s\n", BinaryFormat(0))
	fmt.Printf("   1的二进制是 %s\n", BinaryFormat(1))
	fmt.Printf("  -1的二进制是 %s\n", BinaryFormat(-1))
	fmt.Printf(" 260的二进制是 %s\n", BinaryFormat(260))
	fmt.Printf("-260的二进制是 %s\n", BinaryFormat(-260))

	fmt.Printf("65536的二进制是 %s\n", BinaryFormat(65536))
	fmt.Printf("-65536的二进制是 %s\n", BinaryFormat(-65536))
	fmt.Printf("2184364的二进制是 %s\n", BinaryFormat(2184364))
	fmt.Printf("-2184364的二进制是 %s\n", BinaryFormat(-2184364))
}

// 有实现格式化输出，不过只提交作业内容即可
