package main

import (
	"fmt"
	"math/rand"
)

func sum() {
	const N1 = 8
	const N2 = 4
	A := [N1][N2]float64{}

	for m := 0; m < N1; m++ {
		for n := 0; n < N2; n++ {
			A[m][n] = rand.Float64()
		}

	}
	// fmt.Println(A)

	B := [N1][N2]float64{}

	for m := 0; m < N1; m++ {
		for n := 0; n < N2; n++ {
			B[m][n] = rand.Float64()
		}
	}

	sum := 0.
	for i := 0; i < N1; i++ {
		for j := 0; j < N2; j++ {
			sum += A[i][j] + B[i][j]
			// fmt.Println(sum)
		}
	}
	fmt.Printf("两个矩阵的和为: %f\n", sum)
}

func in_season1(m int) {
	if m > 0 && m <= 3 {
		fmt.Printf("%d月为春季\n", m)
	} else if m > 3 && m <= 6 {
		fmt.Printf("%d月为夏季\n", m)
	} else if m > 6 && m <= 9 {
		fmt.Printf("%d月为秋季\n", m)
	} else if m > 9 && m <= 12 {
		fmt.Printf("%d月为冬季\n", m)
	} else {
		fmt.Println("1-3为春季，4-6为夏季，7-9为秋季，10-12为冬季。")
	}
}

func in_season2(n int) {
	switch n {
	case 1, 2, 3:
		fmt.Printf("%d月为春季\n", n)
	case 4, 5, 6:
		fmt.Printf("%d月为夏季\n", n)
	case 7, 8, 9:
		fmt.Printf("%d月为秋季\n", n)
	case 10, 11, 12:
		fmt.Printf("%d月为冬季\n", n)
	default:
		fmt.Println("1-3为春，4-6为夏，7-9为秋，10-12为冬。")
	}
}

type Name_results struct {
	Name        string
	yw          float32
	mathematics float32
	english     float32
}

var nm Name_results

func (nr Name_results) average() { //每位同学的平均分
	fmt.Printf("%s的平均分是%.2f\n", nr.Name, (nr.english+nr.mathematics+nr.yw)/3)
}

func (_ Name_results) average1(sli []Name_results) { //整个班三门课的平均分
	var sum float32
	for i, _ := range sli {
		// fmt.Println(i, y)
		sum += sli[i].yw + sli[i].english + sli[i].mathematics
	}
	// fmt.Println(sum)
	fmt.Printf("整个班三门课的平均分为%.2f\n", sum/float32(len(sli)*3))
}

func (_ Name_results) average2(sli []Name_results) { //全班同学平均分低于60的
	sum := 0
	for n, _ := range sli {
		// fmt.Println(n, m)
		pjf := (sli[n].mathematics + sli[n].english + sli[n].yw) / 3
		// fmt.Println(pjf)
		if pjf < 60.0 {
			sum += 1
		}

	}
	fmt.Printf("全班同学平均分低于60的有%d位\n", sum)
}

func main() {
	// 1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
	// sum()

	// 2.给定月份，判断属于哪个季节。分别用if和switch实现
	// for i := 0; i <= 13; i++ {
	// 	in_season1(i) //测试数据
	// 	in_season2(i)
	// }

	// 3.创建一个student结构体，包含姓名和语数外三门课的成绩。
	// 用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位

	//添加数据
	sli := make([]Name_results, 0, 10)
	sli = append(sli, Name_results{"tom", 52.1, 60.2, 78.3})
	sli = append(sli, Name_results{"jerry", 58.1, 68.2, 79.3})
	sli = append(sli, Name_results{"xm", 65.3, 78.2, 89.3})
	sli = append(sli, Name_results{"xh", 33.2, 33.2, 33.2})
	sli = append(sli, Name_results{"abc", 28.1, 3.2, 89.2})
	// fmt.Println(sli)

	for _, m := range sli {
		// fmt.Println(n, m)
		m.average() //每位同学的平均分
	}

	nm.average1(sli) //整个班三门课的平均分

	nm.average2(sli) //全班同学平均分低于60的

}

// 第一可以实现一个生成矩阵的方法进行调用
// 第二题在不符合的逻辑处理上要说明为什么不符合
// 第三题可以使用生成随机数
