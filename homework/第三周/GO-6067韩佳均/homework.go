package main

import (
	"fmt"
	"math/rand"
)

func work1() {
	A := [8][5]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			A[i][j] = rand.Intn(100)
		}
	}
	B := [8][5]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			B[i][j] = rand.Intn(100)
		}
	}
	C := [8][5]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}

	}
	fmt.Println("两个矩阵的和为：", C)

}
func work2_if(mouth int) {
	if mouth >= 1 && mouth <= 3 {
		fmt.Println("这是春天")
	} else if mouth >= 4 && mouth <= 6 {
		fmt.Println("这是夏天")
	} else if mouth >= 7 && mouth <= 9 {
		fmt.Println("这是秋天")
	} else if mouth >= 10 && mouth <= 12 {
		fmt.Println("这是冬天")
	} else {
		fmt.Println("输入月份失败")
	}
}
func work2_swich(mouth int) {
	switch {
	case mouth >= 1 && mouth <= 3:
		fmt.Println("这是春天")
	case mouth >= 4 && mouth <= 6:
		fmt.Println("这是夏天")
	case mouth >= 7 && mouth <= 9:
		fmt.Println("这是秋天")
	case mouth >= 10 && mouth <= 12:
		fmt.Println("这是冬天")
	default:
		fmt.Println("输入月份失败")
	}
}

type student struct {
	name    string
	English int
	math    int
	chinese int
}

func work3() {
	// 用一个slice容纳一个班的同学
	class := []student{}
	// Name := [...]int{1, 2, 3}
	name := [...]string{"小韩", "小涂", "小汪", "小李", "小鹏", "小童", "小陈", "小骠"}
	for i := 0; i < len(name); i++ {
		English := rand.Intn(100)
		math := rand.Intn(100)
		chinese := rand.Intn(100)
		class = append(class, student{string(name[i]), English, math, chinese})

	}
	English_sum := 0
	math_sum := 0
	chinese_sum := 0
	count := 0
	for i := 0; i < len(class); i++ {
		avg := (class[i].math + class[i].English + class[i].chinese) / 3
		fmt.Printf("%s同学的平均分是%d\n", class[i].name, avg)
		//低于60分的同学
		if avg < 60 {
			count++
		}
		English_sum += class[i].English
		math_sum += class[i].math
		chinese_sum += class[i].chinese
	}
	fmt.Printf("全班同学平均分低于60的有%d位同学\n", count)
	fmt.Printf("整个班英语的平均分是%d\n", English_sum/(len(name)))
	fmt.Printf("整个班数学的平均分是%d\n", math_sum/(len(name)))
	fmt.Printf("整个班语文的平均分是%d\n", chinese_sum/(len(name)))

}

func main() {
	work1()
	work2_if(7)
	work2_swich(13)
	work3()
}
