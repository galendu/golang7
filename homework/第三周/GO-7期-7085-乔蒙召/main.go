package main

import (
	"fmt"
	"math/rand"
	"time"
)


func SumMatrix(a [8][5]int, b [8][5]int) [8][5]int {
	c := [8][5]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func RandArr() [8][5]int {
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	// num := r.Intn(100)
	arr := [8][5]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			arr[i][j] = rand.Intn(100)
		}
	}
	return arr
}

func PrintArr(arr [8][5]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%3d ", arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

type month = int
func GetSeasonIf(m month) {
	if m <= 0 {
		fmt.Printf("month error!\n")
	} else if m <= 3 {
		fmt.Printf("Spring.\n")
	} else if m <= 6 {
		fmt.Printf("Summer.\n")
	} else if m <= 9 {
		fmt.Printf("Autumn.\n")
	} else if m <= 12 {
		fmt.Printf("Winter.\n")
	} else {
		fmt.Printf("month error!\n")
	}
}

func GetSeasonSwitch1(m month){
	switch m {
	case 1:
		fmt.Printf("Spring.\n")
	case 2:
		fmt.Printf("Spring.\n")
	case 3:
		fmt.Printf("Spring.\n")
	case 4:
		fmt.Printf("Summer.\n")
	case 5:
		fmt.Printf("Summer.\n")
	case 6:
		fmt.Printf("Summer.\n")
	case 7:
		fmt.Printf("Autumn.\n")
	case 8:
		fmt.Printf("Autumn.\n")
	case 9:
		fmt.Printf("Autumn.\n")
	case 10:
		fmt.Printf("Winter.\n")
	case 11:
		fmt.Printf("Winter.\n")
	case 12:
		fmt.Printf("Winter.\n")
	default:
		fmt.Printf("month error!\n")
	}
}

func GetSeasonSwitch2(m month){
	switch {
	case m<=0:
		fmt.Printf("month error!\n")
	case m<=3:
		fmt.Printf("Spring.\n")
	case m<=6:
		fmt.Printf("Summer.\n")
	case m<=9:
		fmt.Printf("Autumn.\n")
	case m<=12:
		fmt.Printf("Winter.\n")
	default:
		fmt.Printf("month error!\n")
	}
}

type student struct {
	name string
	chinese int
	math int
	english int
	avg float32
}

//初始化班级成员及成绩信息
func InitSlice() []student{
	// s:="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s:="ABCDEFG" //HIJKLMNOPQRSTUVWXYZ"
	stusinfo := make([]student, 0, len(s))
	for i:=0; i<len(s); i++{
		stu_name:=string(s[i])
		stu_chinese:=rand.Intn(100+1)
		stu_math:=rand.Intn(100+1)
		stu_english:=rand.Intn(100+1)
		stu_avg:=float32(stu_chinese+stu_math+stu_english)/3
		info:=student{name:stu_name, chinese: stu_chinese, math: stu_math, english: stu_english, avg: stu_avg}
		stusinfo = append(stusinfo, info)
	}
	return stusinfo
}

//统计班级成员信息
type classinfo struct{
	stusinfo []student
	avg_chinese float32
	avg_math float32
	avg_english float32
	avgnopass int
}
func StatStusinfo(stusinfo []student){
	var total_chinese, total_math, total_english int
	var avgnopass int
	for _,info := range stusinfo{
		total_chinese += info.chinese
		total_math += info.math
		total_english += info.english
		if info.avg < 60{
			avgnopass++
		}
	}
	avg_chinese:=float32(total_chinese)/float32(len(stusinfo))
	avg_math:=float32(total_math)/float32(len(stusinfo))
	avg_english:=float32(total_english)/float32(len(stusinfo))
	clsinfo:=classinfo{stusinfo: stusinfo, avg_chinese: avg_chinese, avg_math: avg_math, avg_english: avg_english, avgnopass: avgnopass}

	for _, info:=range clsinfo.stusinfo{
		fmt.Printf("name: %s, chinese: %d, math: %d, english: %d, avg: %.2f\n", info.name, info.chinese, info.math, info.english, info.avg)
	}
	fmt.Printf("avg_chinese: %.2f, avg_math: %.2f, avg_english: %.2f, avgnopass: %d\n", clsinfo.avg_chinese, clsinfo.avg_math, clsinfo.avg_english, clsinfo.avgnopass)
}

func main() {
	//1
	fmt.Printf("作业1\n")
	rand.Seed(time.Now().Unix()) //生成随机种子如果放到RandArr里，会导致随机的 a b 数组完全相同 ？
	a := RandArr()
	b := RandArr()
	c := SumMatrix(a, b)
	PrintArr(a)
	PrintArr(b)
	PrintArr(c)
	fmt.Printf("-----------------------------\n")

	//2
	fmt.Printf("作业2\n")
	m := rand.Intn(12) + 1
	fmt.Printf("month: %d\n", m)
	GetSeasonIf(m)
	GetSeasonSwitch1(m)
	GetSeasonSwitch2(m)
	fmt.Printf("-----------------------------\n")

	//3
	fmt.Printf("作业3\n")
	stusinfo:=InitSlice()
	fmt.Printf("班级成员及成绩信息\n")
	for _, info := range stusinfo{
		fmt.Printf("name: %s, chinese: %d, math: %d, english: %d\n", info.name, info.chinese, info.math, info.english)
	}
	fmt.Println()
	StatStusinfo(stusinfo)
	fmt.Printf("-----------------------------\n")
}
