package main

import (
	"fmt"
	"math"
	"strings"
)
//��������Ƶ�һ������
func BinaryFormat(n int32) string {
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31))
	for i := 0; i < 32; i++ {
		if a&c != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}

		c >>= 1

	}
	return sb.String()

}
func main() {
	fmt.Printf("(8)   ���������: %s \n", BinaryFormat(8))
	fmt.Printf("(120)  ���������: %s \n", BinaryFormat(260))
	fmt.Printf("(-280)  ���������: %s \n", BinaryFormat(-280))
	fmt.Printf("(1024)  ���������: %s \n", BinaryFormat(1024))
	fmt.Printf("(-20968) ���������: %s \n", BinaryFormat(-20968))

}