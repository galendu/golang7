package main

import (
	"fmt"
	"math/rand"
	"time"
)

func zy1() {
	sli := make([]int, 0, 10)
	for i := 1; i <= 100; i++ {
		rand.Seed(time.Now().UnixNano())
		sli = append(sli, rand.Intn(128))
	}
	fmt.Printf("sli 共有%d个元素\n%v\n", len(sli), sli)

	var m = make(map[int]int)
	for k, y := range sli {
		m[y] = k
	}

	fmt.Printf("map中共有%d个互不相同的元素\n%v\n", len(m), m)
}

func main() {
    //1. 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次
    //往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
	zy1()

    //2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}
    //返回“2 4 6”。输入的切片可能很短，也可能很长。
    zy2()

}