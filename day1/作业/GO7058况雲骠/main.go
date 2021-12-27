package main

import (
	"firstwoke/Binary"
	"fmt"
)

func main() {
	var a int32 = 0
	var b int32 = 1
	var c int32 = -1
	var d int32 = 260
	var e int32 = -260
	fmt.Printf("0 %s\n", Binary.BinaryFormat(a))
	fmt.Printf("1 %s\n", Binary.BinaryFormat(b))
	fmt.Printf("-1 %s\n", Binary.BinaryFormat(c))
	fmt.Printf("260 %s\n", Binary.BinaryFormat(d))
	fmt.Printf("-260 %s\n", Binary.BinaryFormat(e))
}
