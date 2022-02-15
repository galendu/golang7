package d

import (
	"go-course/package/a/b/c/internal"
	"go-course/package/a/b/c/internal/e/f"
)

var D int

func Foo() {
	internal.Internal = 4
	f.F = 4
}
