package main

import (
	"fmt"
	"math"
	"strings"
)

func BinaryFormat(p int32) string {
        s := strings.Builder{}
        c := uint32(math.Pow(2, 31))
        a := uint32(p)
        for i := 0; i < 32; i++ {
                if c&a != 0 {
                        s.WriteString("1")
                } else {
                        s.WriteString("0")
                }
                c >>= 1
        }
        return s.String()
}

func main() {
        var (
                p1 int32 = 0
                p2 int32 = 1
                p3 int32 = -1
                p4 int32 = 260
                p5 int32 = -260
        )
        fmt.Println(BinaryFormat(p1))
        fmt.Println(BinaryFormat(p2))
        fmt.Println(BinaryFormat(p3))
        fmt.Println(BinaryFormat(p4))
        fmt.Println(BinaryFormat(p5))
}


// 建议输出时格式化输出，还要注意代码缩进
