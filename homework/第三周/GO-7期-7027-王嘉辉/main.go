package main

import (
	"fmt"
	"math/rand"
)

func add111() {
	var sum int = 0
	const L1 = 3
	const L2 = 4
	A := [L1][L2]int{}
	B := [L1][L2]int{}
	for a := 0; a < L1; a++ {
		for b := 0; b < L2; b++ {
			A[a][b] = rand.Intn(100)
		}
	}

	for a := 0; a < L1; a++ {
		for b := 0; b < L2; b++ {
			B[a][b] = rand.Intn(100)
		}
	}
	for a := 0; a < L1; a++ {
		for b := 0; b < L2; b++ {
			sum = sum + A[a][b] + B[a][b]
		}
	}
	fmt.Println(sum)
}

func season111(m int) {
	fmt.Printf("这个月是%d月\n", m)
	if m > 0 && m <= 3 {
		fmt.Printf("%d月是春季\n", m)
	} else if m > 3 && m <= 6 {
		fmt.Printf("%d月是夏季\n", m)
	} else if m > 6 && m <= 9 {
		fmt.Printf("%d月是秋季\n", m)
	} else if m > 3 && m <= 6 {
		fmt.Printf("%d月是冬季\n", m)
	} else {
		fmt.Println("请输入正确的月分")
	}
}

func season222(m int) {
	fmt.Printf("这个月是%d月\n", m)
	switch m {
	case 1, 2, 3:
		fmt.Printf("%d月是春季\n", m)
	case 4, 5, 6:
		fmt.Printf("%d月是夏季\n", m)
	case 7, 8, 9:
		fmt.Printf("%d月是秋季\n", m)
	case 10, 11, 12:
		fmt.Printf("%d月是冬季\n", m)
	default:
		fmt.Println("请输入正确的月分")
	}
}

type student struct {
	name    string
	chinese int
	math    int
	english int
}

func student111() {
	crr := []student{}
	name := []string{"张三", "李四", "王五"}
	for i := 0; i < len(name); i++ {
		english := rand.Intn(100)
		math := rand.Intn(100)
		chinese := rand.Intn(100)
		crr = append(crr, student{name[i], chinese, math, english})
	}
	fmt.Println(crr)
	sum := 0 //小于60分的次数
	avg := 0 //全班所有科目平均分
	englishavg := 0
	mathavg := 0
	chineseavg := 0
	for i := 0; i < len(name); i++ {
		avgscore := (crr[i].english + crr[i].math + crr[i].chinese) / len(name)
		fmt.Printf("%s同学的平均分是%d\n", crr[i].name, avgscore)
		avg += avgscore
		englishavg += crr[i].english / len(name)
		mathavg += crr[i].math / len(name)
		chineseavg += crr[i].chinese / len(name)
		if avgscore < 60 {
			sum += 1
		}
	}
	fmt.Printf("整个班平均分数小于60的有%d个\n", sum)
	fmt.Printf("语文的平均分是%d\n", chineseavg)
	fmt.Printf("数学的平均分是%d\n", mathavg)
	fmt.Printf("英语的平均分是%d\n", englishavg)
	fmt.Printf("全班同学的所有科目平均分是%d\n", avg/len(name))
}

func main() {
	add111()
	season111(4)
	// fmt.Println(" ")
	season222(9)
	student111()
}

// 第一题可以再理解下题意，整体逻辑是正确的。第二题在整理下是不是少了一个判断逻辑