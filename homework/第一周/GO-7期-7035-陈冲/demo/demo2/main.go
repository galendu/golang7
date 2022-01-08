package main

// 导入包
import (
	"fmt"

	"gonum.org/v1/gonum/stat" // 引入第三方包
)

// 全局变量
var (
	username = "xiaoming"
	address  = "beijing"
	sex      string
	age      int
	weight   float64
)

// 定义常量
const (
	PI    = 3.14
	LEVEL = "000000"
)

// 变量
func variable_demo() {
	fmt.Printf("username: %s\n", username)
	username := "Tom" // 函数区域内定义局部变量, 也叫内部作用域，就近原则，打印最近的一个变量值
	fmt.Printf("username: %s\n", username)

	fmt.Printf("address: %s\n", address)
	fmt.Printf("sex: %s\n", sex)         // 默认赋值空字符
	fmt.Printf("age: %d\n", age)         // 默认 0
	fmt.Printf("weight: %.2f\n", weight) // 默认 0.00  这里保留了2位
}

// 常量
func const_demo() {
	fmt.Printf("%f\n", PI)
	// LEVEL = "111111"  // 不能更改常量
	fmt.Printf("%s\n", LEVEL)

	// 执行结果
	// 3.140000
	// 000000
}

// iota 是一个特殊常量值，是系统定义的可以被编译器修改的常量值。
// iota 只能被用在常量的赋值中，在每一个const关键字出现时，被重置为0，然后每出现一个常量，
// iota 所代表的数值会自动增加1，iota可以理解成常量组中常量的计数器，不论该常量的值是什么，只要有一个常量，那么iota 就加1。
func iota_demo() {
	const (
		a = iota
		b
		c
		d
	)

	fmt.Printf("%d, %d, %d, %d\n", a, b, c, d)
	// 0, 1, 2, 3
	const (
		a1 = iota
		b1
		c1
		d1 = 10
		f1 // 和上一行相同
		f2 // 和上一行相同
		e1 = iota
		e2
	)

	fmt.Printf("%d, %d, %d, %d, %d, %d, %d, %d\n", a1, b1, c1, d1, f1, f2, e1, e2)
	// 0, 1, 2, 10, 10, 10, 6, 7

	const (
		// 计算存储容量  1024 = 2^10  => 1TB = 2^10GB = 2^20MB = 2^30KB
		_  = iota             // iota =0
		TB = 1 << (10 * iota) // iota =1
		GB = 1 << (10 * iota) // iota =2
		MB = 1 << (10 * iota) // iota =3
		KB = 1 << (10 * iota) // iota =4
	)
	fmt.Printf("KB=%d, MB=%d, GB=%d, TB=%d\n", KB, MB, GB, TB)
	// KB=1099511627776, MB=1073741824, GB=1048576, TB=1024
}

// 指针
func pointer() {
	var a int
	p := &a // p用%p和%d输出的结果是一样的，只是形式不一样（十六进制和十进制）
	fmt.Printf("%p %p\n", p, &a)
	fmt.Printf("%d 0x%x\n", p, p)

	// 0xc000015c88 0xc000015c88
	// 824633810056 0xc000015c88
}

//literal 字面量
func literal() {
	fmt.Printf("%t\n", 04 == 4.00)      // 用到了一个整型字面量和浮点型字面量 // true
	fmt.Printf("%v\n", .4i)             // 虚数字面量 0.4i
	fmt.Printf("%t\n", '\u4f17' == '众') // Unicode和rune字面量   // true
	fmt.Printf("%c\n", '众')             // 众 %c 表示一个字符
	fmt.Printf("%c\n", '\u4f17')        // 众
	fmt.Printf("%c\n", '\u263A')        // Unicode 笑脸 ☺
	fmt.Printf("Hello\nWorld\n!\n")     //字符串字面量
}

func main() {
	x := []float64{1, 2, 3, 4, 5, 6} // 定义一个类型为float64的数组，并初始化
	// 计算 x 的方差
	vari := stat.Variance(x, nil)
	fmt.Printf("方差是 %v\n", vari)

	variable_demo()
	const_demo()
	iota_demo()
	pointer()
	literal()
}

/*
1、这里执行报错
~/go/src/magedu.com.com/demo2 » go run main.go                                                                                                      wadechen@MacBookPro
main.go:7:2: no required module provides package gonum.org/v1/gonum/stat; to add it:
        go get gonum.org/v1/gonum/stat
--------------------------------------------------------------------------------------------------------------------------------------------------------------------
使用 go get gonum.org/v1/gonum/stat 安装
~/go/src/magedu.com.com/demo2 » go get gonum.org/v1/gonum/stat                                                                                  1 ↵ wadechen@MacBookPro
go get: added gonum.org/v1/gonum v0.9.3
--------------------------------------------------------------------------------------------------------------------------------------------------------------------
在此执行 go run main.go
~/go/src/magedu.com.com/demo2 » go run main.go                                                                                                      wadechen@MacBookPro
方差是 3.5
*/
