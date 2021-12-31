package util

import (
	"math"
	"strings"
)

// DecimalToBinary 输入一个32位整数的二进制表示，包括正负数
func DecimalToBinary(num int32) string {
	var result strings.Builder
	var value = uint32(num)
	cursor := uint32(math.Pow(2, 31)) // 不能强转为int32，会导致溢出
	for i := 0; i < 32; i++ {
		temp := value & cursor
		if temp == 0 {
			result.WriteString("0")
		} else {
			result.WriteString("1")
		}
		cursor >>= 1
	}
	return result.String()
}
