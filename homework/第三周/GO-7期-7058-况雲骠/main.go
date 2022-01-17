package main

import (
	"fmt"
	"math/rand"
)

type student struct {
	chinese float32
	math    float32
	english float32
	name    byte
}

func work1() {
	fmt.Println("作业1")
	array1 := [8][5]int{}
	array2 := [8][5]int{}
	array3 := [8][5]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			array1[i][j] = rand.Intn(100)
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			array2[i][j] = rand.Intn(100)
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			array3[i][j] = array1[i][j] + array2[i][j]
			fmt.Printf("%d\t", array3[i][j])
		}
		fmt.Printf("\n")
	}

}

func work2() {
	fmt.Println("作业2")
	month := rand.Intn(12)
	fmt.Printf("%d月\n", month)
	//用if实现
	fmt.Println("if:")
	if month >= 10 {
		fmt.Println("这是冬天")
	} else if month >= 7 {
		fmt.Println("这是秋天")
	} else if month >= 4 {
		fmt.Println("这是夏天")
	} else {
		fmt.Println("这是春天")
	}

	//用switch实现
	fmt.Println("switch：")
	switch {
	case month >= 10:
		fmt.Println("这是冬天")
	case month >= 7:
		fmt.Println("这是秋天")
	case month >= 4:
		fmt.Println("这是夏天")
	default:
		fmt.Println("这是春天")
	}

}

func work3() {
	fmt.Println("作业3")
	var s [20]student = [20]student{}
	S := make([]student, 20)

	for i, name := 0, 'A'; i < 20; i, name = i+1, int32(name)+1 {
		s[i].name = byte(name)
		s[i].chinese = rand.Float32() * 100
		s[i].math = rand.Float32() * 100
		s[i].english = rand.Float32() * 100
		S[i].name = s[i].name
		S[i].chinese = s[i].chinese
		S[i].math = s[i].math
		S[i].english = s[i].english
		fmt.Printf("学生%d的姓名是%c\n语文是%f\n数学是%f\n英语是%f\n", i+1, s[i].name, s[i].chinese, s[i].math, s[i].english)
	}
	var avagec, avagem, avagee, avagea float32 = 0.0, 0.0, 0.0, 0.0 //各项分数
	num := 0                                                        //不及格人数

	for i := range S {
		avagec = avagec + S[i].chinese
		avagem = avagem + S[i].math
		avagee = avagee + S[i].english
		avagea = avagec + avagem + avagee
	}
	avagec /= 20
	avagem /= 20
	avagee /= 20
	avagea /= 60
	for i := range S {
		if score := S[i].chinese + S[i].math + S[i].english; score < 180.0 {
			num++
		}
	}
	fmt.Printf("全科平均成绩%f\n语文平均成绩%f\n数学平均成绩%f\n英语平均成绩%f\n综合成绩挂科人数%d\n", avagea, avagec, avagem, avagee, num)

}
func main() {
	work1()
	work2()
	work3()

}
