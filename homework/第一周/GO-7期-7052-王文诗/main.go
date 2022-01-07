package main

import (
	"fmt"
)

func f1(n int32) {
	fmt.Printf("%.32b\n", n)
}

func main() {
	f1(0)
	f1(1)
	f1(-1)
	f1(260)
	f1(-260)
}

// 作业按照作业要求来
