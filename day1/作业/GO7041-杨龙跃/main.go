package main

import (
	"fmt"
	"math"
	"strings"
)

//输出一个int32对应的二进制表示
func BinaryFormat(n int32) string {
	a := uint32(n)//将传进来的int32类型数值转换为无符号的uint32并赋值给a
	sb := strings.Builder{}//拼接字符串
	c := uint32(math.Pow(2, 31))//求2的31次方的幂值 二进制为：10000000000000000000000000000000
	for i := 0; i < 32; i++ {
		if a&c > 0 {
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
	fmt.Println(BinaryFormat(-1))
	fmt.Println(BinaryFormat(1))
	fmt.Println(BinaryFormat(260))
	fmt.Println(BinaryFormat(-260))

}

