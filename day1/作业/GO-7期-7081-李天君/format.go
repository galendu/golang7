package main
import (
	"fmt"
	"math"
	"strings"
)

func BinaryFormat(n int32) string{
	a := uint32(n)
	sb := strings.Builder{}
	// 第32位是1，后面1~31全是0 ， 100000000...
	c := uint32(math.Pow(2, 31))
	for i := 0; i < 32 ; i++ {
		if  a&c != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1
	}
	return sb.String()
}

func main()  {
	//format()
	fmt.Printf("5的二进制:%s\n",BinaryFormat(5))
	fmt.Printf("-5的二进制:%s\n",BinaryFormat(-5))
	fmt.Printf("7的二进制:%s\n",BinaryFormat(7))
	fmt.Printf("-7的二进制:%s\n",BinaryFormat(-7))
	fmt.Printf("177的二进制:%s\n",BinaryFormat(177))
	fmt.Printf("6888877的二进制:%s\n",BinaryFormat(6888877))
}