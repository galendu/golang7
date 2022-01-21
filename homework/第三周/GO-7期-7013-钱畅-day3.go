//day3
package main

import (
	"fmt"
	"math/rand"
)

func one(row int, col int) {
	const rowsize int = 5
	const colsize int = 8
	arr1 := [rowsize][colsize]int{}
	for i := 0; i < rowsize; i++ {
		for j := 0; j < colsize; j++ {
			arr1[i][j] = rand.Intn(64)
			fmt.Printf("%d ", arr1[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
	arr2 := [rowsize][colsize]int{}
	for i := 0; i < rowsize; i++ {
		for j := 0; j < colsize; j++ {
			arr2[i][j] = rand.Intn(64)
			fmt.Printf("%d ", arr2[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
	//矩阵逐元素相加
	arr3 := [rowsize][colsize]int{}
	for i := 0; i < rowsize; i++ {
		for j := 0; j < colsize; j++ {
			arr3[i][j] = arr1[i][j] + arr2[i][j]
			fmt.Printf("%d ", arr3[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}

func two1(n int) {
	month := n
	if month >= 1 && month <= 3 {
		fmt.Printf("%d"+" month is spring", month)
	} else if month >= 4 && month <= 6 {
		fmt.Printf("%d"+" month is summer", month)
	} else if month >= 7 && month <= 9 {
		fmt.Printf("%d"+" month is autumn", month)
	} else if month >= 10 && month <= 12 {
		fmt.Printf("%d"+" month is winter", month)
	} else {
		fmt.Print("no this month")
	}
	fmt.Println()
}

func two2(m int) {
	month := m
	switch month {
	case 1, 2, 3:
		fmt.Printf("%d"+" month is spring", month)
	case 4, 5, 6:
		fmt.Printf("%d"+" month is summer", month)
	case 7, 8, 9:
		fmt.Printf("%d"+" month is autumn", month)
	case 10, 11, 12:
		fmt.Printf("%d"+" month is winter", month)
	default:
		fmt.Print("no this month")
	}
	fmt.Println()
}

type student struct {
	name                string
	lang, english, math float64
	// english float32
	// math    float32
}

func three(stu []student) (map[string]float64, map[string]float64) {
	//return (int(stu.lang)+int(stu.englist)+int(stu.math))/3
	mapstu := make(map[string]float64)
	mapsub := make(map[string]float64)
	sumlang, sumenglish, summath := 0., 0., 0.
	for index, value := range stu {
		mapstu[value.name] = float64((value.lang + value.english + value.math) / 3)
		sumlang += value.lang
		sumenglish += value.english
		summath += value.math
		mapsub["langavg"] = sumlang / float64(index+1)
		mapsub["englishavg"] = sumenglish / float64(index+1)
		mapsub["mathavg"] = summath / float64(index+1)
	}
	//println("%d", stuavg)
	return mapstu, mapsub
}

func main() {
	one(8, 5)
	two1(10)
	two2(5)
	slicestu := []student{{"yi", 20, 40, 90}, {"er", 20, 50, 100}, {"san", 10, 20, 80}, {"si", 30, 70, 90}, {"wu", 20, 50, 70}, {"liu", 15, 35, 95}}
	m1, m2 := three(slicestu)
	for key, value := range m1 {
		fmt.Printf("%v=%v\t", key, value)
	}
	fmt.Println()
	for key, value := range m2 {
		fmt.Printf("%v=%v\t", key, value)
	}
}
