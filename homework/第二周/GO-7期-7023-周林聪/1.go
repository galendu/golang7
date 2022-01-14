package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//要求：创建一个初始长度为0、容量为10的切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不同的元素
func foo() {
	arr := make([]int, 0, 100)
	map1 := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(128))
		key := arr[i]                                //将切片里的元素赋值为map1的key
		if _, exists := map1[key]; exists == false { //如果这个key第一次出现
			map1[key] = 1 //将此key对应的value作为计数作用，value从1开始计数
		} else { //如果这个key重复出现
			map1[key] += 1 //重复出现一次，就统计一次
		}
	}
	fmt.Println(arr)
	fmt.Println(map1)
	fmt.Println(len(map1)) //因为map的key不会重复，所以map1的len就是互不同元素的个数

	//打印结果:
	// [33 15 71 59 1 6 57 44 72 36 70 47 34 113 88 26 11 21 37 98 15 90 104 18 127 43 47 120 54 119 53 120 91 15 37 76 41 119 125 18 13 18 74 67 113 19 94 100 127 89 21 73 117 23 40 17 72 122 103 43 3 30 61 28 106 36 105 2 31 34 75 104 106 118 103 38 73 7 124 52 31 121 1 29 9 43 105 19 3 50 62 88 2 115 36 7 88 35 13
	// 126]
	// map[1:2 2:2 3:2 6:1 7:2 9:1 11:1 13:2 15:3 17:1 18:3 19:2 21:2 23:1 26:1 28:1 29:1 30:1 31:2 33:1 34:2 35:1 36:3 37:2 38:1 40:1 41:1 43:3 44:1 47:2 50:1 52:1 53:1 54:1 57:1 59:1 61:1 62:1 67:1 70:1 71:1 72:2 73:2 74:1 75:1 76:1 88:3 89:1 90:1 91:1 94:1 98:1 100:1 103:2 104:2 105:2 106:2 113:2 115:1 117:1 118:1
	// 119:2 120:2 121:1 122:1 124:1 125:1 126:1 127:2]
	// 69
}

//要求：实现一个函数func arr2string(arr []int)string,比如输入[]int{2 4 6}，返回“2 4 6”。输入的切片可能很短，也可能很长
func arr2string(arr []int) string {
	sb := strings.Builder{}
	for i := 0; i < len(arr); i++ {
		value := arr[i]
		if i < len(arr)-1 { //这样避免最后多出一个空格
			sb.WriteString(fmt.Sprint(strconv.Itoa(value))) //将int转成string类型
			sb.WriteString(" ")                             //元素之间添加空格
		} else {
			sb.WriteString(fmt.Sprint(strconv.Itoa(value))) //最后一个元素
		}
	}
	sb1 := sb.String()
	return sb1
}

func main() {
	foo()

	arr1 := []int{}
	arr2 := []int{2, 4, 6, 8, 10}
	arr3 := []int{2, 4}

	fmt.Printf("arr2string(arr1)==%v\n", arr2string(arr1))
	fmt.Printf("arr2string(arr2)==%v\n", arr2string(arr2))
	fmt.Printf("arr2string(arr3)==%v\n", arr2string(arr3))
	fmt.Printf("arr2string(nil)==%v\n", arr2string(nil))

	//打印结果
	// arr2string(arr1)==
	// arr2string(arr2)==2 4 6 8 10
	// arr2string(arr3)==2 4
	// arr2string(nil)==
}

// 只输出要求的结果即可