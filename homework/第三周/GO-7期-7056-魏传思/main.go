package main

import (
	"fmt"
	"math/rand"
)

// 作业1
// 随机初始化两个8*5的矩阵，求两个矩阵的和
func mat_add() {
	const N = 8
	const M = 5
	A := [N][M]int{}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			A[i][j] = rand.Intn(100)
		}
	}
	B := [N][M]int{}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			B[i][j] = rand.Intn(100)
		}
	}
	C := [N][M]int{}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			C[i][j] = A[i][j] + B[i][j]

		}
	}
	fmt.Println(A)
	fmt.Println()

	fmt.Println(B)
	fmt.Println()

	fmt.Println(C)
}

// 作业2
// 给定月份，判断属于哪个季节。分别用if 和 switch实现
// if 实现
func est_season() {
	var i int = 2
	if i > 0 && i <= 3 {
		fmt.Printf("%d月属于春季", i)
	} else if i > 3 && i <= 6 {
		fmt.Printf("%d月属于夏季", i)
	} else if i > 6 && i <= 9 {
		fmt.Printf("%d月属于秋季", i)
	} else if i > 9 && i <= 12 {
		fmt.Printf("%d月属于冬季", i)
	} else {
		fmt.Printf("不存在%d月", i)
	}
}

// switch 实现
func est_season2() {
	var n int = 13
	switch {
	case (0 < n && n <= 3):
		fmt.Printf("%d月属于春季", n)
	case (3 < n && n <= 6):
		fmt.Printf("%d月属于夏季", n)
	case (6 < n && n <= 9):
		fmt.Printf("%d月属于秋季", n)
	case (9 < n && n <= 12):
		fmt.Printf("%d月属于冬季", n)
	default:
		fmt.Printf("不存在%d月", n)
	}
}

// 作业3
// 创建一个student结构体，包含姓名和语数外三门课的成绩。
// 用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位

type student struct {
	name         string
	chineseScore int
	mathScore    int
	foreignScore int
}

func aa(n int) []student {
	st := make([]student, n)
	for i := 0; i < n; i++ {
		stu := student{"name", rand.Intn(100), rand.Intn(100), rand.Intn(100)}
		st = append(st, stu)
	}
	return st

}

func (s student) culScore(st []student) {
	for k, v := range st {
		alAve := (v.chineseScore + v.foreignScore + v.mathScore) / 3
		fmt.Printf("第%d位同学的平均分是%d\n", k, alAve)
		sum := 0
		if alAve < 60 {
			sum++
			fmt.Printf("全班同学的平均分低于60的有%d位\n", sum)
		}
	}

}

func main() {
	// mat_add()

	// est_season()
	// fmt.Println()

	// est_season2()
	// fmt.Println()

	//第三题代码没全做出来
	stu := student{}
	st := aa(100)
	stu.culScore(st)

}
