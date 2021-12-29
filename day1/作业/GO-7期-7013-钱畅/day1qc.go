//作业:输出一个int32对应的二进制表示

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
		if a/c >= 1 {
			sb.WriteString("1")
			a -= c
		} else {
			sb.WriteString("0")
		}
		c /= 2
	}
	return sb.String()
}

func main() {
	num := BinaryFormat(21)
	fmt.Printf(num)
}
