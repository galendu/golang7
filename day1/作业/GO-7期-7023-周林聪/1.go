//要求:输出一个int32对应的2进制表示
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
		if a&c != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c = c >> 1
	}
	return sb.String()
}

func main() {
	var aaa uint32 = 260
	var bbb uint32 = 660

	fmt.Printf(" 260 == %032b\n", 260)
	fmt.Printf(" 260 == %s\n", BinaryFormat(260))
	fmt.Printf("-260 == %b\n", -aaa)
	fmt.Printf("-260 == %s\n", BinaryFormat(-260))

	fmt.Println(BinaryFormat(0))

	fmt.Printf(" 660 == %032b\n", 660)
	fmt.Printf(" 660 == %s\n", BinaryFormat(660))
	fmt.Printf("-660 == %b\n", -bbb)
	fmt.Printf("-660 == %s\n", BinaryFormat(-660))
}

//go run homework/oneday/1.go
//测试。。。。
