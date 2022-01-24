//这里是关于包的简单说明
package main

import (
	"fmt"
	"go-course/entrance_class/util"
	"runtime"
	"strconv"
)

//arithmetic 算术运算
func arithmetic() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = a + b
	var d float32 = a - b
	var e float32 = a * b
	var f float32 = a / b
	fmt.Printf("a=%.3f, b=%.3f, c=%.3f, d=%.3f, e=%.3f, f=%.3f\n", a, b, c, d, e, f)
}

//relational 关系运算符
func relational() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = 8
	fmt.Printf("a==b吗 %t\n", a == b)
	fmt.Printf("a!=b吗 %t\n", a != b)
	fmt.Printf("a>b吗 %t\n", a > b)
	fmt.Printf("a>=b吗 %t\n", a >= b)
	fmt.Printf("a<c吗 %t\n", a < b)
	fmt.Printf("a<=c吗 %t\n", a <= c)
}

//logistic 逻辑运算符
func logistic() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = 8
	fmt.Printf("a>b && b>c吗 %t\n", a > b && b > c)
	fmt.Printf("a>b || b>c吗 %t\n", a > b || b > c)
	fmt.Printf("a>b不成立，对吗 %t\n", !(a > b))
	fmt.Printf("b>c不成立，对吗 %t\n", !(b > c))
}

//bit_op 位运算
func bit_op() {
	fmt.Printf("os arch %s, int size %d\n", runtime.GOARCH, strconv.IntSize) //int是4字节还是8字节，取决于操作系统是32位还是64位
	var a int32 = 260
	fmt.Printf("260     %s\n", util.BinaryFormat(a))
	fmt.Printf("-260    %s\n", util.BinaryFormat(-a)) //负数用补码表示。在对应正数二进制表示的基础上，按拉取反，再末位加1
	fmt.Printf("260&4   %s\n", util.BinaryFormat(a&4))
	fmt.Printf("260|3   %s\n", util.BinaryFormat(a|3))
	fmt.Printf("260^7   %s\n", util.BinaryFormat(a^7))     //^作为二元运算符时表示异或
	fmt.Printf("^-260   %s\n", util.BinaryFormat(^-a))     //^作为一元运算符时表示按位取反，符号位也跟着变
	fmt.Printf("-260>>10 %s\n", util.BinaryFormat(-a>>10)) //正数高位补0，负数高位补1
	fmt.Printf("-260<<3 %s\n", util.BinaryFormat(-a<<3))   //负数左移，可能变成正数
	//go语言没有循环（无符号）左/右移符号   >>>  <<<
}

//assignment 赋值运算
func assignment() {
	var a, b int = 8, 3
	a += b
	fmt.Printf("a+=b %d\n", a)
	a, b = 8, 3
	a -= b
	fmt.Printf("a-=b %d\n", a)
	a, b = 8, 3
	a *= b
	fmt.Printf("a*=b %d\n", a)
	a, b = 8, 3
	a /= b
	fmt.Printf("a/=b %d\n", a)
	a, b = 8, 3
	a %= b
	fmt.Printf("a%%=b %d\n", a) //%在fmt里有特殊含意，所以需要前面再加个%转义一下
	a, b = 8, 3
	a <<= b
	fmt.Printf("a<<=b %d\n", a)
	a, b = 8, 3
	a >>= b
	fmt.Printf("a>>=b %d\n", a)
	a, b = 8, 3
	a &= b
	fmt.Printf("a&=b %d\n", a)
	a, b = 8, 3
	a |= b
	fmt.Printf("a|=b %d\n", a)
	a, b = 8, 3
	a ^= b
	fmt.Printf("a^=b %d\n", a)
}

func main() {
	arithmetic()
	fmt.Println()
	relational()
	fmt.Println()
	logistic()
	fmt.Println()
	bit_op()
	fmt.Println()
	assignment()
}

//go run entrance_class/operator/main.go
