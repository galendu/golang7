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
	c := uint32(math.Pow(2, 31)) //最高位上是1，其他位是0
	for i := 0; i < 32; i++ {
		if a&c != 0 { //
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1
	}
	return sb.String()
}

func main() {

	fmt.Println(BinaryFormat(128))
	fmt.Println(BinaryFormat(-128))
	fmt.Println(BinaryFormat(1088))

}
