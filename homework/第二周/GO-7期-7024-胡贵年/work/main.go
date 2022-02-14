package main

// 1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
// 2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。

import (
	"fmt"
	"math/rand"
	"strings"
)

//作业1
func statistics() {
	//定义一个切片变量
	s1 := make([]int, 0, 10)

	//生成随机数，并存入 s1 切片
	for i := 0; i < 100; i++ {
		s1 = append(s1, rand.Intn(128))
	}

	// 定义一个map
	var m map[int]int
	m = make(map[int]int)

	// 把切片的数据存入map的key里
	for _, ele := range s1 {
		//赋值给map， 并计数
		m[ele] += 1
	}

	// 打印map的key值
	for key, vlaue := range m {
		fmt.Printf("%d:%d ", key, vlaue)
	}

	// 统计map的长度，得到切片里不同元素的个数
	fmt.Printf("\n统计：%d\n", len(m))
}

// 作业2
func arr2string(arr []int) string {

	// 定义一个大的String
	sb := strings.Builder{}

	// 往String里添加数据
	for _, ele := range arr {
		sb.WriteString(fmt.Sprint(ele))
		sb.WriteString(" ")
	}

	// 返回String
	merged := sb.String()
	return merged

	// return ""
}

func main() {
	// 作业1
	statistics()

	fmt.Println("==============================================")

	// 作业2
	arr1 := []int{2, 4, 6, 7, 8, 2}
	v := arr2string(arr1)
	fmt.Println(v)

}
