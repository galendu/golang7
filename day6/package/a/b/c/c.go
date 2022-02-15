package c

import (
	"go-course/package/a/b/c/internal"
	"go-course/package/a/b/c/internal/e/f"
)

var C int

func Foo() {
	internal.Internal = 4
	f.F = 4
}
