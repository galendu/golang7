package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Student struct {
	Name         string
	ScopeChinese int
	ScopeMath    int
	ScopeEnglish int
}

func main() {
	all_student := make([]Student, 40, 40)
	//初始化
	for i := 0; i < 40; i++ {
		all_student[i].Name = "刘" + strconv.Itoa(i)
		all_student[i].ScopeChinese = rand.Intn(100)
		all_student[i].ScopeEnglish = rand.Intn(100)
		all_student[i].ScopeMath = rand.Intn(100)
	}

	//计算科目总和
	var sum_chinese int
	var sum_english int
	var sum_math int

	//计算每位的平均分和学科总和
	var avg_evenybody = make(map[string]float32, 40)
	for _, ele := range all_student {
		avg_evenybody[ele.Name] = float32((ele.ScopeMath + ele.ScopeEnglish + ele.ScopeChinese) / 3)
		sum_chinese += ele.ScopeChinese
		sum_english += ele.ScopeEnglish
		sum_math += ele.ScopeMath
	}
	//遍历输出成绩
	var flag int
	for k, v := range avg_evenybody {
		fmt.Printf("%s的平均分为%.1f\n", k, v)
		if v < 60 {
			flag += 1
		}
	}
	fmt.Println("========================")

	//计算三门学科的平均分
	fmt.Printf("Math的平均分为%.1f\n", float32(sum_math/len(all_student)))
	fmt.Printf("English的平均分为%.1f\n", float32(sum_english/len(all_student)))
	fmt.Printf("Chinese的平均分为%.1f\n", float32(sum_chinese/len(all_student)))

	//计算全班均分低于60的人数
	fmt.Println("========================")
	fmt.Printf("低于60分的同学有%d位", flag)

}
