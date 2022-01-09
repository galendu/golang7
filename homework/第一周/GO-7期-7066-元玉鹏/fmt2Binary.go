package BinaryFormat

import (
	"strings"
)

//Prints a string of binary expressions for a number of type INT32
func BinaryFormat(n int32) string {

	// Consider the plus and minus signs
	unsignNum := uint32(n)
	strBuilder := strings.Builder{}
	caliper := (^uint32(0) >> 1) + 1

	for caliper > 0 {
		if 0 == caliper&unsignNum {
			strBuilder.WriteString("0")
		} else {
			strBuilder.WriteString("1")
		}
		caliper >>= 1
	}

	return strBuilder.String()
}

// func main() {
// 	fmt.Println("1 -> ", BinaryFormat(1))
// 	fmt.Println("2 -> ", BinaryFormat(2))
// 	fmt.Println("-1 -> ", BinaryFormat(-1))
// 	fmt.Println("235 -> ", BinaryFormat(235))

// }

/*
	注意：此函数输出的二进制表达形式为原码形式，实际的加减运算逻辑都是在补码的形势下进行计算的
*/

// 单文件建议使用package main
