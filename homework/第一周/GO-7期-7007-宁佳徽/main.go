package main

import (
	"fmt"
	"math"
	"strings"
)

func BinaryFormat(n int32) string {
	// 将传进来的参数n赋值给a
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31))
	for i := 0; i < 32; i++ {
		// 使用a 和c这个2^31次方进行 按位与运算，如果结果不等于0，则说明当前位上的值为1，反之则为0
		if a&c != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		// c右移1位后赋值给c，一共循环32次
		c >>= 1
	}
	return sb.String()
}
func main() {
	var e int32 = 300
	fmt.Printf("e=%s\n", BinaryFormat(e))
	fmt.Printf("e=%s\n", BinaryFormat(-e))

}

//逻辑是正确，建议完成作业的内容
