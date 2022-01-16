package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

/*
1.1 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素
1.2 利用map统计该切片里有多少个互不相同的元素。
*/

func NewSc(number int) (sc []int) {
	sc = make([]int, 0, 10)
	for num := 1; num <= number; num ++ {
		sc = append(sc, rand.Intn(128))
	}
	return sc
}

func ElementsArchive(sc []int) (cm map[int]int) {

	cm = make(map[int]int, 10)
	for _, el := range sc {
		if _, ok := cm[el]; ok {
			cm[el] ++
		} else {
			cm[el] = 1
		}
	}
	return cm
}

func ElementsNumber(cm map[int]int) (num int) {
	return len(cm)
}


// 2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。

func ArrToString(arr []int) string {
	var tmp string

	for _, value := range arr {
		tmp = fmt.Sprintf("%s %s", tmp, strconv.Itoa(value))
	}
	return strings.TrimSpace(tmp)
}


func main()  {
	nsc := NewSc(100)
	cea := ElementsArchive(nsc)

	// 1
	fmt.Printf("The number of different elements is :: %d\n", ElementsNumber(cea))
	// 2
	fmt.Printf("The number of different elements is :: %d\n", len(cea))


	atsArr1 := []int{1, 23, 4, 5}
	fmt.Printf("%#v %T\n", ArrToString(atsArr1), ArrToString(atsArr1))

	//
	atsArr2 := []int{1, 2, 4, 5, 0, 9, 1, 2, 3, 4, 5, 8}
	fmt.Printf("Slice To Strings:: %#v %T\n", ArrToString(atsArr2), ArrToString(atsArr2))
}
