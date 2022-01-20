// 2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func arr2string(arr []int) string {
	sb := strings.Builder{} // 使用这种方式拼接字符串
	for _, ele := range arr {
		sb.WriteString(strconv.Itoa(ele))
		sb.WriteString(" ")
	}
	return strings.Trim(sb.String(), " ") // 去除末尾空格
}

func main() {
	s := []int{2, 4, 6}
	fmt.Printf("result: [%s]\n", arr2string(s))
}

