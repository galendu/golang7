package main

import (
	"strings"
	"math"
	"fmt"
)

func BinaryFormat(n int32) string{
	a:=uint32(n)
	sb:=strings.Builder{}
	c:=uint32(math.Pow(2,31))

	for i:=0; i<32;i++{
		if a&c > 0 {
			sb.WriteString("1")
		}else{
			sb.WriteString("0")
		}
		c >>= 1 
	}
	return sb.String()

}

func main(){

	fmt.Printf("%032b\n",260)
	fmt.Printf("0=%s\n",BinaryFormat(0))
	fmt.Printf("1=%s\n",BinaryFormat(1))
	fmt.Printf("-1=%s\n",BinaryFormat(-1))
	fmt.Printf("260=%s\n",BinaryFormat(260))
	fmt.Printf("-260=%s\n",BinaryFormat(-260))
}
