package main

import (
	"fmt"
	"math"
	"strings"
)

// 输出一个int32对应的的二进制表示
func BinaryFormat(n int32) string {
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31))

	// a和c进行与运算
	//
	for i := 0; i < 32; i++ {
		// 进行与运算，只有a和 c的二进制在同一位置上都为1时，得到的值不等于0，其他情况得到的值都是0
		if a&c != 0 {
			// 不等于0时输出1
			sb.WriteString("1")
		} else {
			// 等于0时输出0
			sb.WriteString("0")
		}
		//  c 右移一位，之后再和a做与运算
		c >>= 1
	}
	return sb.String()

}

func main() {
	// 正数打印原码、负数打印补码
	fmt.Println("12434")
	fmt.Printf("%f\n", 15.3)
	fmt.Printf("%s\n", BinaryFormat(-5))
	fmt.Printf("%b\n", uint32(math.Pow(2, 31)))
	fmt.Printf("%s\n", BinaryFormat(5))
	fmt.Printf("%s\n", BinaryFormat(-2147483648)) //极值
	fmt.Printf("%s\n", BinaryFormat(2147483647))  // 2147483648报错
	fmt.Printf("%s\n", BinaryFormat(-0))
	fmt.Printf("%s\n", BinaryFormat(0))
	// fmt.Printf("%s\n",BinaryFormat(1456343452))
	// fmt.Printf("%s\n",BinaryFormat(-563457586))

}
