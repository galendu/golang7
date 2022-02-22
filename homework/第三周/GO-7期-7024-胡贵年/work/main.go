package main

import (
	"fmt"
	"math/rand"
)

// 1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
// 2.给定月份，判断属于哪个季节。分别用if和switch实现
// 3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位

//作业1
func arr_sum(arr1 [8][5]int, arr2 [8][5]int) int {

	// 初始化数组
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[0]); j++ {
			arr1[i][j] = rand.Intn(10)
			fmt.Printf("%d ", arr1[i][j])
		}
		fmt.Println()
	}
	fmt.Println("==================")
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[0]); j++ {
			arr2[i][j] = rand.Intn(10)
			fmt.Printf("%d ", arr1[i][j])
		}
		fmt.Println()
	}

	// 求和
	sum := 0
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[0]); j++ {
			sum += arr1[i][j] + arr2[i][j]
		}
	}

	// 返回sum
	return sum
}

// 作业2
func judge_season(number int) string {

	var season string
	switch number {
	case 1, 2, 3:
		season = "春天"
	case 4, 5, 6:
		season = "夏天"
	case 7, 8, 9:
		season = "秋天"
	case 10, 11, 12:
		season = "冬天"
	default:
		fmt.Println("输入的月份有误，请重新输入[ 1-12 ]")

	}
	return season
}

// 作业3
// 3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
type Class struct {
	class_name string
	stu        []Student
	avg_eng    float32
	avg_lan    float32
	avg_math   float32
}
type Student struct {
	cla_name       string
	name           string
	sex            byte
	language_Score int
	math_Score     int
	english_Score  int
}

func (stu Student) stu_avg_score() float32 {
	return ((float32)(stu.english_Score+stu.language_Score+stu.math_Score) / 3)
}

func (cl Class) avg_score(cla *Class) {
	var eng_sum float32
	var lan_sum float32
	var math_sum float32
	for i := 0; i < len(cla.stu); i++ {
		eng_sum += float32(cla.stu[i].english_Score)
		lan_sum += float32(cla.stu[i].language_Score)
		math_sum += float32(cla.stu[i].math_Score)
	}
	cla.avg_eng = eng_sum / float32(len(cla.stu))
	cla.avg_lan = lan_sum / float32(len(cla.stu))
	cla.avg_math = math_sum / float32(len(cla.stu))

}

func (cla Class) lower_score() []Student {
	lower_stu := make([]Student, 0, 10)
	for i := 0; i < len(cla.stu); i++ {
		if cla.stu[i].stu_avg_score() < 60 {
			lower_stu = append(lower_stu, cla.stu[i])
		}
	}
	return lower_stu
}

func main() {
	// 作业1
	fmt.Println("==================== 作业1 =======================")
	arr1 := [8][5]int{}
	arr2 := [8][5]int{}
	sum := arr_sum(arr1, arr2)
	fmt.Printf("sum=%d\n", sum)

	// 作业2
	fmt.Println()
	fmt.Println("==================== 作业2 =======================")
	fmt.Println()
	number := 51
	season := judge_season(number)
	fmt.Printf("%s\n", season)

	// 作业3
	fmt.Println()
	fmt.Println("==================== 作业3 =======================")
	fmt.Println()
	stu01 := Student{}
	stu01.cla_name = "北京2班"
	stu01.name = "包打听"
	stu01.sex = 1
	stu01.language_Score = rand.Intn(81) + 20
	stu01.math_Score = rand.Intn(81) + 20
	stu01.english_Score = rand.Intn(81) + 20

	stu02 := Student{}
	stu02.cla_name = "北京2班"
	stu02.name = "顺风耳"
	stu02.sex = 1
	stu02.language_Score = rand.Intn(81) + 20
	stu02.math_Score = rand.Intn(81) + 20
	stu02.english_Score = rand.Intn(81) + 20

	stu03 := Student{}
	stu03.cla_name = "北京2班"
	stu03.name = "顺风耳"
	stu03.sex = 1
	stu03.language_Score = rand.Intn(81) + 20
	stu03.math_Score = rand.Intn(81) + 20
	stu03.english_Score = rand.Intn(81) + 20

	stu04 := Student{}
	stu04.cla_name = "北京2班"
	stu04.name = "二郎神"
	stu04.sex = 1
	stu04.language_Score = rand.Intn(81) + 20
	stu04.math_Score = rand.Intn(81) + 20
	stu04.english_Score = rand.Intn(81) + 20

	stu05 := Student{}
	stu05.cla_name = "北京2班"
	stu05.name = "吉吉国王"
	stu05.sex = 1
	stu05.language_Score = rand.Intn(81) + 20
	stu05.math_Score = rand.Intn(81) + 20
	stu05.english_Score = rand.Intn(81) + 20

	cla := Class{}
	cla.class_name = "北京2班"
	cla.stu = append(cla.stu, stu01, stu02, stu03, stu04, stu05)

	// 计算每位同学的平均分
	for i := 0; i < len(cla.stu); i++ {
		fmt.Printf("来自 %s 的 %s 同学 语文：%d , 数学：%d , 英语：%d , 平均分是 %f \n", cla.stu[i].cla_name, cla.stu[i].name, cla.stu[i].language_Score, cla.stu[i].math_Score, cla.stu[i].english_Score, cla.stu[i].stu_avg_score())
	}
	// 计算班级各科的平均分
	cla.avg_score(&cla)
	fmt.Printf("%s 的英语平均分：%f 语文的平均分：%f 数学的平均分：%f \n", cla.class_name, cla.avg_eng, cla.avg_lan, cla.avg_math)

	// 统计这个班级低于60分的同学
	low_score := cla.lower_score()
	for _, stu := range low_score {
		fmt.Printf("同学 %s 平均分低于60分,为 %f\n", stu.name, stu.stu_avg_score())
	}
}
