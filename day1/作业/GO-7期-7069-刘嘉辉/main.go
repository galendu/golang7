package main

import "fmt"

func main() {
	fmt.Printf("0的二进制是： %s\n", BinaryFormat(0))
	fmt.Printf("1的二进制是： %s\n", BinaryFormat(1))
	fmt.Printf("-1的二进制是： %s\n", BinaryFormat(-1))
	fmt.Printf("260的二进制是： %s\n", BinaryFormat(260))
	fmt.Printf("-260的二进制是： %s\n", BinaryFormat(-260))
}
