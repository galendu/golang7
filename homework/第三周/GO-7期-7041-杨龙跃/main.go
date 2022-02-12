package main

import (
	"fmt"
	"math/rand"
)

/*
1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）
2.给定月份，判断属于哪个季节。分别用if和switch实现
3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，
求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位.
*/
func homework1() {
	//随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）

	arr1 := [8][5]int{}
	arr2 := [8][5]int{}
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 2; j++ {
			arr1[i][j] = rand.Intn(66)
			arr2[i][j] = rand.Intn(66)
			arr3 := arr1[i][j] + arr2[i][j]
			//求两个矩阵的和
			fmt.Println("arr3:", arr3)

		}

	}

}

func homework2_1(moth int) {
	//给定月份，判断属于哪个季节。if实现
	if moth >= 1 && moth <= 3 {
		fmt.Printf("输入的月份%v是春季\n", moth)
	} else if moth > 3 && moth <= 6 {
		fmt.Printf("输入的月份%v是夏季\n", moth)
	} else if moth > 6 && moth <= 9 {
		fmt.Printf("输入的月份%v是秋季\n", moth)
	} else if moth > 9 && moth <= 12 {
		fmt.Printf("输入的月份%v是冬季\n", moth)
	} else {
		fmt.Println("输入的月份有误!!!")
	}

}

func homework2_2(moth int) {
	//给定月份，判断属于哪个季节。switch实现
	switch {
	case moth >= 1 && moth <= 3:
		fmt.Printf("输入的月份%v是春季\n", moth)

	case moth > 3 && moth <= 6:
		fmt.Printf("输入的月份%v是夏季\n", moth)

	case moth > 6 && moth <= 9:
		fmt.Printf("输入的月份%v是秋季\n", moth)
	case moth > 9 && moth <= 12:
		fmt.Printf("输入的月份%v是冬季\n", moth)

	}

}

func homework3() {
	/*
			创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，
		求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
	*/
	//创建一个student结构体，包含姓名和语数外三门课的成绩
	type student struct {
		name   string
		yuwen  float64
		shuxue float64
		waiyu  float64
	}

	//一个slice容纳一个班的同学
	var students = make([]student, 0)
	students = append(students, student{"张一", 70.0, 65.7, 23.9})
	students = append(students, student{"张二", 70.0, 65.7, 12.9})
	students = append(students, student{"张三", 45.0, 55.7, 33.9})
	students = append(students, student{"张四", 70.0, 65.7, 60.9})
	students = append(students, student{"张五", 70.0, 65.7, 88.9})
	fmt.Println("students：", students)
	count := 0.0
	var lower_stu = make([]string, 0)
	for _, v := range students {
		// fmt.Println("v:",v)
		//求每位同学的平均分
		fmt.Printf("同学%v的三门成绩的平均成绩：%.2f\n", v.name, (v.shuxue+v.waiyu+v.yuwen)/3)
		if (v.shuxue+v.waiyu+v.yuwen)/3 < 60 {
			lower_stu = append(lower_stu, v.name)
		}
		scorce := v.shuxue + v.waiyu + v.yuwen
		count += scorce
	}
	//整个班三门课的平均分
	fmt.Printf("整个班三门课的总成绩%f,学生%v个，平均分:%.2f\n", count, len(students), count/float64(len(students)))
	//全班同学平均分低于60的有几位
	fmt.Printf("全班同学平均分低于60有%v位,他们的名字是%v\n", len(lower_stu), lower_stu)

}

func main() {
	homework2_1(11)
	homework2_2(11)
	homework3()
	homework1()

}

// 第一题可以实现下生成矩阵的方法
// 第二题为啥switch方法少了一个逻辑
