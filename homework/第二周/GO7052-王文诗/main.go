package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var s1 = make([]int, 0, 10)
var m1 = make(map[int]int, 100)

/*
创建一个初始长度为0，容量为10的int型切片，
调用rand.Intn(128)100次，往切片里边添加100个元素，
利用map统计该切片里有多个互不相同的元素
*/
func sliceAdd() {
	// 从0~100循环100次
	for i := 0; i < 100; i++ {
		// 随机生成0~128之间的数字
		rand.Seed(time.Now().UnixNano())
		r1 := rand.Intn(128)
		// 添加到int类型的切片中
		s1 = append(s1, r1)
		// 添加到map中做去重操作
		m1[r1] = r1
	}
	fmt.Println(s1)
	fmt.Println(len(m1))
	// 上边随机生成0~128之间的数字，会有一个问题，就是每次循环随机生成的数字列表都是一样的，解决办法是加一个rand.Seed(time.Now().UnixNano())
}

/*
实现一个函数slcString(slc []int) string
输入int型切片，返回一个string类型的字符串
*/

// slcString 函数接收一个int类型的切片，返回一个string类型的字符串
func slcString(slc []int) string {
	// 定义一个string类型的切片并且初始化
	s1 := make([]string, 0)
	// 判断传入的slc是空直接返回空字符串
	if len(slc) > 0 {
		// 循环传入的slc
		for _, i := range slc {
			// 循环把slc切片中的值全部转换成字符串
			v := strconv.Itoa(i)
			// 写入到s1切片中
			s1 = append(s1, v)
		}
	}
	// 传入一个string类型的切片和分隔符，返回一个string
	result := strings.Join(s1, " ")
	return result
}

func main() {
	sliceAdd()
	//input := make([]int,0)
	//input = []int{1,2,3,4,5,6,7,8,3,25,2,12,45,67,64,2,45,2,5,2,42}
	//slcS:=slcString(input)
	//fmt.Printf("type:%T value:%#v\n",slcS,slcS)
}

// 建议格式化输出
// 第二题要考虑处理输出结果末尾的空格