package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

// 1.随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）

// JudMonth 2.给定月份，判断属于哪个季节。分别用if和switch实现
func JudMonth(month int) (season string, err error) {
	switch {
	case month > 0 && month <= 3:
		season = "Spring"
	case month >= 4 && month <= 6:
		season = "Summer"
	case month >= 7 && month <= 9:
		season = "Autumn"
	case month >= 10 && month <= 12:
		season = "Winter"
	default:
		season = "null"
	}

	if season == "null" {
		return "", errors.New(fmt.Sprintf("month need in [1..12]"))
	}
	return fmt.Sprintf("%s", season), nil
}

// 3.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
type Student struct {
	UserName    string
	Language    int
	Mathematics int
	English     int
}

func NewStudent(username string, lang, math, eng int) (student *Student) {
	return &Student{
		UserName:    username,
		Language:    lang,
		Mathematics: math,
		English:     eng,
	}
}

func (stu *Student) ComputeSumStudent() (css int) {
	return stu.Language + stu.Mathematics + stu.English
}

func (stu *Student) ComputeAvgStudent() (cas int) {
	return stu.ComputeSumStudent() / 3
}

func MakeOneClass() (class []*Student) {
	class = make([]*Student, 0)
	stus := fmt.Sprintf("%s%s", MakeKeyword(), strings.ToUpper(MakeKeyword()))
	for i := 0; i < len(stus); i++ {
		class = append(class, NewStudent(string(stus[i]), rand.Intn(100)+1, rand.Intn(100)+1, rand.Intn(100)+1))
	}
	return class
}

func MakeKeyword() (keyword string) {

	for i := 'a'; i < 'z'; i++ {
		keyword = fmt.Sprintf("%s%c", keyword, i)
	}
	return keyword
}

func main() {
	jm, err := JudMonth(1)
	if err != nil {
		fmt.Println("err::", err)
	} else {
		fmt.Println(jm)
	}

	moc := MakeOneClass()
	var onenum, totalnum int
	for _, v := range moc {
		// 每个人的
		onenum = v.ComputeAvgStudent()
		fmt.Printf("stuend %s avg :: %d \n", v.UserName, onenum)
		// 整个班级
		totalnum = totalnum + v.ComputeSumStudent()
	}
	fmt.Printf("stuends avg :: %d \n", totalnum/len(moc))

}

