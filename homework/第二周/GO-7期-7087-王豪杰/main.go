package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

//type IntSlice []int
//
//func (s IntSlice) Len() int  {
//	return len(s)
//}
//
//func (s IntSlice) Swap(i,j int)  {
//	s[i],s[j] = s[j],s[i]
//}
//
//func (s IntSlice) Less(i,j int) bool  {
//	return s[i] < s[j]
//}

// 随机数
func randomNumber(arr []int) []int {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(128))
	}
	return arr
}

//统计不重复个数
func statisticsNumber(arr []int) int {
	list := make(map[int]bool, 100)
	for _, slice := range arr {
		list[slice] = true
	}
	return len(list)
}

// 返回字符拼接
func sliceString(arr []int) string {
	sb := strings.Builder{}
	//if len(arr) == 0 {
	//	return "nil"
	//}
	for _, slice := range arr {
		sb.WriteString(strconv.Itoa(slice))
		sb.WriteString(" ")
	}
	return sb.String()
}

func main() {
	// 随机数
	Slice := make([]int, 0, 100)
	result := randomNumber(Slice)
	sort.IntSlice.Sort(result)
	fmt.Println(result)

	//统计不重复个数
	result1 := statisticsNumber(result)
	fmt.Println(result1)

	// 返回字符拼接
	arr := make([]int, 0, 100)
	result2 := sliceString(arr)
	fmt.Println(result2)
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result2 = sliceString(arr)
	fmt.Println(result2)

}

// 第一题输出格式化下，建议使用判断map中key是否存在，输出要求即可
// 第二题考虑处理下输出结果末尾的空格
