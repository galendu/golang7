package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// 1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
func slice_demo() {
	slice1 := make([]int, 0, 10)
	map1 := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		slice1 = append(slice1, rand.Intn(128))
		key := slice1[i]
		if _, value := map1[key]; value == false {
			map1[key] = 1 // 没有就添加
		} else {
			map1[key] += 1 // 重复就加1
		}
	}

	fmt.Printf("切片：%v\n", slice1)
	fmt.Printf("map: %v\n", map1)
	fmt.Printf("切片中不通的元素个数：%d\n", len(map1))

	// 切片：[33 15 71 59 1 6 57 44 72 36 70 47 34 113 88 26 11 21 37 98 15 90 104 18 127 43 47 120 54 119 53 120 91 15 37 76 41 119 125 18 13 18 74 67 113 19 94 100 127 89 21 73 117 23 40 17 72 122 103 43 3 30 61 28 106 36 105 2 31 34 75 104 106 118 103 38 73 7 124 52 31 121 1 29 9 43 105 19 3 50 62 88 2 115 36 7 88 35 13 126]
	// map: map[1:2 2:2 3:2 6:1 7:2 9:1 11:1 13:2 15:3 17:1 18:3 19:2 21:2 23:1 26:1 28:1 29:1 30:1 31:2 33:1 34:2 35:1 36:3 37:2 38:1 40:1 41:1 43:3 44:1 47:2 50:1 52:1 53:1 54:1 57:1 59:1 61:1 62:1 67:1 70:1 71:1 72:2 73:2 74:1 75:1 76:1 88:3 89:1 90:1 91:1 94:1 98:1 100:1 103:2 104:2 105:2 106:2 113:2 115:1 117:1 118:1 119:2 120:2 121:1 122:1 124:1 125:1 126:1 127:2]
	// 切片中不通的元素个数：69
}

// 2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
func arr2string(arr []int) string {
	strb := strings.Builder{}
	for i := 0; i < len(arr); i++ {
		value := arr[i]
		if i < len(arr)-1 {
			strb.WriteString(fmt.Sprint(strconv.Itoa(value))) //  strconv.Itoa 将int 转 string
			strb.WriteString(" ")
		} else { // 如果是最后一个元素
			strb.WriteString(fmt.Sprint(strconv.Itoa(value)))
		}
	}

	return strb.String()
}

func main() {
	slice_demo()

	arr1 := []int{}
	arr2 := []int{1}
	arr3 := []int{2, 3, 4}

	fmt.Printf("arr1转换结果：%v\n", arr2string(arr1))
	fmt.Printf("arr2转换结果：%v\n", arr2string(arr2))
	fmt.Printf("arr3转换结果：%v\n", arr2string(arr3))

	// arr1转换结果：
	// arr2转换结果：1
	// arr3转换结果：2 3 4
}

// 整体完成的不错。只输出作业要求即可
