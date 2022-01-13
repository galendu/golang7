/*
实现一个函数func arr2string(arr []int) string，
比如输入[]int{2,4,6}，返回“2 4 6”。
输入的切片可能很短，也可能很长。
*/
package homework2

import (
	"fmt"
	"strconv"
	"strings"
)

//func strings.Join(elems []string, sep string) string
func Arr2string(arr []int) string {
	str := []string{}
	for _, ele := range arr {
		//必须从切片中取出元素做强制转换
		str = append(str, fmt.Sprint(strconv.Itoa(ele)))
	}
	return strings.Join(str, " ")
}

//输出
