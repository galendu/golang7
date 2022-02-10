//作业1，把一个int32的数换算成二进制
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(BinaryFormat(0))
	fmt.Println(BinaryFormat(-1))
	fmt.Println(BinaryFormat(1))
	fmt.Println(BinaryFormat(260))
	fmt.Println(BinaryFormat(-260))
}

//输出一个int32对应的二进制表示
func BinaryFormat(n int32) string {
	a := uint32(n)               //将int32类型的数值转换成无符号的uint32类型
	sb := strings.Builder{}      //拼接字符串的函数
	c := uint32(math.Pow(2, 31)) //二进制为：10000000 00000000 00000000 00000000
	for i := 0; i < 32; i++ {
		if a&c > 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1 //每次右偏移量1和n的二进制进行比对
		//fmt.Printf("当前二进制是%b<<<<<<<<<<\n",c)
	}
	fmt.Printf("%d的二进制是%b,%d的十进制是%d\n", n, a, n, n)
	fmt.Printf("2的31次方的二进制是%b,十进制是%d\n", c, c)

	return "换算成32进制是" + sb.String()
}
