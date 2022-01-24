package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//创建一个初始长度为0,容量为0的int型切片，调用rand.Intn(128)100次，随机向切片中添加100个元素。
func a01() []int {
	arr := make([]int, 0, 100)   //len=0,cap=100
	rand.Seed(time.Now().Unix()) //构建一个种子,用来每次执行的时候生成不同的切片
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(128))
	}
	return arr
}

//统计map切片中多少个不同的元素
func a2() map[int]int {
	b1 := a01()

	b2 := make(map[int]int)
	for i := 0; i < len(b1); i++ {
		count := 0
		for j := 0; j < len(b1); j++ {
			if b1[i] == b1[j] { //计算slice02中相同元素的个数
				count++
			}
		}
		b2[b1[i]] = count //将slice02中的元素作为map中的key，统计个数作为map中的value进行存放到wcmap中
	}
	return b2
}

func arr2string(arr []int) string {
	str := strings.Builder{} //调用字符串拼接
	for i := 0; i < len(arr); i++ {
		str.WriteString(strconv.Itoa(arr[i])) //将整形变量转换为字符串
		str.WriteString(" ")
	}
	s := str.String() //将拼接之后的字符转换为字符串
	return s
}

func main() {
	a3 := a2()
	fmt.Println(a3)
	fmt.Printf("该切片中有%d个互不相同的元素\n", len(a3))

	arr := []int{2, 4, 6}
	fmt.Println(arr2string(arr))

}
