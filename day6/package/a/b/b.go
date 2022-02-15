package main

import (
	"fmt"
	// "go-course/package/a/b/c/internal"     //go build时会报错：use of internal package go-course/a/b/c/internal not allowed
	// "go-course/package/a/b/c/internal/e/f" //go build时会报错：use of internal package go-course/a/b/c/internal/e/f not allowed
)

var B int

// func f1() {
// 	internal.Internal = 4
// }

// func f2() {
// 	f.F = 4
// }

func main() {
	// f1()
	// f2()
	fmt.Printf("B=%d\n", B)
}

//go build -o b go-course/package/a/b
