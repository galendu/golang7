package main

import (
	"go-course/package/a/b/c"
	"go-course/package/a/b/c/d"
)

var A int

func main() {
	c.Foo()
	d.Foo()
}

//go build -o a go-course/package/a
