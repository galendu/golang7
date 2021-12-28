package utils

import (
	"math"
	"strings"
)

/*
BinaryFormat 函数点睛之笔
第一次循环：比如入计算5的二进制
c => 1000 000 000 000 000 000 000 000 000 000
a => 0000 000 000 000 000 000 000 000 000 101
r => 0000 000 000 000 000 000 000 000 000 000 => 0 当前最高位为0
右移一位
c => 0100 000 000 000 000 000 000 000 000 000
以此类推，末尾第3次循环
c => 0000 000 000 000 000 000 000 000 000 100
a => 0000 000 000 000 000 000 000 000 000 101
r => 0000 000 000 000 000 000 000 000 000 100 => 100 > 0 当前位为1
...
最终结果: 0000 000 000 000 000 000 000 000 000 101
*/
func BinaryFormat(n int32) string {
	a := uint32(n)
	strb := strings.Builder{}
	c := uint32(math.Pow(2, 31)) // 最高位为1，其余为0  1000 000 000 000 000 000 000 000 000 000
	for i := 0; i < 32; i++ {
		// 判断 n 的当前最高位上是否为1
		if a&c != 0 {
			strb.WriteString("1")
		} else {
			strb.WriteString("0")
		}
		c >>= 1 // "1"右移一位
	}
	return strb.String()
}
