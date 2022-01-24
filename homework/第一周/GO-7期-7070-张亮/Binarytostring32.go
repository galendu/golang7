package main

import (
	"fmt"
	"math"
	"strings"
)

func BinarytoString(n int32) string {
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31))

	for i := 0; i < 32; i++ {
		if a&c >= 1 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c = c >> 1
	}
	return sb.String()
}

func main() {
	c := BinarytoString(260)
	fmt.Printf("%s\n", c)
	d := BinarytoString(-260)
	fmt.Printf("%s\n", d)
	e := BinarytoString(1)
	fmt.Printf("%s\n", e)
	f := BinarytoString(-1)
	fmt.Printf("%s\n", f)
	g := BinarytoString(0)
	fmt.Printf("%s\n", g)
	/*
		00000000000000000000000100000100
		11111111111111111111111011111100
		00000000000000000000000000000001
		11111111111111111111111111111111
		00000000000000000000000000000000
	*/

}
