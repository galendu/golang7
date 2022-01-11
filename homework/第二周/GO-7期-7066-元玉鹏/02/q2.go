// 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
package q2

import (
	"fmt"
	"strconv"
	"strings"
)

func arr2string(arr []int, sep string) string {

	// Check whether it is empty
	if arr == nil {
		return ""
	}

	strSlice := make([]string, 0, len(arr))

	for _, value := range arr {
		strSlice = append(strSlice, strconv.Itoa(value))
	}

	return strings.Join(strSlice, sep)
}

// Public arr2string
func Array2String(arr []int, sep string) string {
	fmt.Println(arr2string(arr, sep))
	return arr2string(arr, sep)
}
