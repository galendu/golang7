package format

import (
	"math"
	"strconv"
    "strings"
)

func BinaryFormat(n int) string {
	var sb strings.Builder
	size := strconv.IntSize
	t := int(math.Pow(2, float64(size)-1))
	for i := 1; i <= size; i++ {
		if n&t != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		n <<= 1
	}
	return sb.String()
}
