package main

import (
	"fmt"
	"math"
	"strings"
)

func BinaryFormat(n int32) string { //定义函数BinaryFormat 传入int 输出string
	a := uint32(n)               //强制类型转换成uint32(n)  无符号 赋值给a
	sb := strings.Builder{}      //定义变量sb，用来拼接字符
	c := uint32(math.Pow(2, 31)) //定义变量c 为2^31  无符号 1000..00 31个0
	for i := 0; i < 32; i++ {
		if a&c > 0 { //a与c相与  最开始c的第一位是1 a的第一位如果是1的话 则输出1 否则输出0 写入string.Builder，然后c往右移，与a的再次与 a的第二位如果是1的话 则输出1 否则输出0 写入string.Builder
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c = c >> 1 //将c右移1位 依次与a的第一位 第二位 ... 第三十二位
	}
	return sb.String() //返回拼接后的结果  即为a的二进制
}

func main() {
	fmt.Println(BinaryFormat(2)) //将值传入定义好的函数   测试结果
	fmt.Println(BinaryFormat(20))
	fmt.Println(BinaryFormat(40))
	fmt.Println(BinaryFormat(-2))
}


// 逻辑正确，可考虑格式化输出，只提交作业内容即可