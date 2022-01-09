package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

/*
作业1：
1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
作业2：
2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
*/

func homework1() {
	//作业1：创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
	//创建一个初始化为0，容量为10的切片，并初始化
	var slice1 = make([]int, 0, 10)
	// fmt.Printf("slice1的长度：%d,容量%d,类型是：%T\n", len(slice1), cap(slice1), slice1)cl

	//调用rand.Intn(128)一百次，并将产生的随机数追加到切片中
	for i := 0; i < 100; i++ {
		slice1 = append(slice1, rand.Intn(128))
		// fmt.Printf("第%d次打印rand.Intn(128),获取rand值：%d\n", i, rand.Intn(128))

	}
	fmt.Println(slice1) //此时slice1内部存在100个元素
	// fmt.Printf("slice1的长度:%d,容量:%d\n",len(slice1),cap(slice1))
	//利用map统计该切片中有多少互不相同的元素,大致思路，利用map的key唯一性，将切片中得元素分别作为map的key存入一个空map中，自动去除相同的元素
	//创建一个map,并初始化

	var map1 = make(map[int]int, 100)
	for i, v := range slice1 {
		fmt.Printf("slice的第%d个元素是%d\n", i, v)
		map1[v] = i
	}
	for i, _ := range map1 {
		fmt.Println(i)

	}
	fmt.Println("map1", len(map1))

}

func arr2string(arr []int) string {

	arrstr := make([]string, len(arr))
	//作业2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
	for _, v := range arr {
		s := strconv.Itoa(v) //将int类型数组元素转换为string类型
		arrstr = append(arrstr, s)
	}
	// fmt.Println(arrstr)
	arrstr1 := strings.Join(arrstr, " ")                  //以空格拼接转换后的字符串数组为一个整体的字符串
	arrstr1 = strings.Replace(arrstr1, " ", "", len(arr)) //去除新字符串的空格
	// fmt.Println(arrstr1)

	return "\"" + arrstr1 + "\""
}

func main() {
	homework1()
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("===========")
	array1 := arr2string(a)
	fmt.Println(array1)
	// fmt.Printf("\"%v\"\n",array1)
	fmt.Println("=====")
	fmt.Printf("array1:%T\n", array1)

}
