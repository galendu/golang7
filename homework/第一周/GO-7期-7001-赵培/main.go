package main

import (
	"fmt"
	"math"
	"strings"
)

////输出一个int32对应的二进制表示
func BinaryFormat(n int32) string {
	a := uint32(n) //uint32(b) 将接收到的n值转换为无符号的int32类型并赋值给变量a
	// fmt.Printf("%T\n", a) //uint32
	//定义一个变量sb，用来进行接收所有字符串拼接的字符
	//string.Builder 通过使用一个内部的 slice 来存储数据片段。当开发者调用写入方法的时候，数据实际上是被追加（append）到了其内部的 slice 上。
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31)) //求2的31次方的幂值 二进制为：10000000000000000000000000000000
	for i := 0; i < 32; i++ {
		if a&c > 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1 //等价于c = c >> 1  右移1位
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
// 结果是正确的，在打印的时候可以格式化输出下，比如打印出-1对应的二进制:11111111111111111111111111111111