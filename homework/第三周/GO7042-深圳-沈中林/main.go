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
	switch month {
	case 1, 2, 3:
		season = "Spring"
	case 4, 5, 6:
		season = "Summer"
	case 7, 8, 9:
		season = "Autumn"
	case 10, 11, 12:
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
	// jud month
	jm, err := JudMonth(1)
	if err != nil {
		fmt.Println("err::", err)
	} else {
		fmt.Println(jm)
	}

	moc := MakeOneClass()
	var onenum, langtotal, mathtotal, engtotal int
	var unlod int
	for _, v := range moc {
		// student
		onenum = v.ComputeAvgStudent()
		if onenum < 60 {
			unlod = unlod + 1
		}
		// student avg
		fmt.Printf("stuend %s avg :: %d \n", v.UserName, onenum)

		// stuends
		langtotal = langtotal + v.Language
		mathtotal = mathtotal + v.Mathematics
		engtotal = engtotal + v.English
	}

	// stuends < 60
	fmt.Printf("stuends < 60 :: %d \n", unlod)
	// students avg for lang
	fmt.Printf("stuends Language avg :: %d \n", langtotal/len(moc))
	// students avg for math
	fmt.Printf("stuends Mathematics avg :: %d \n", mathtotal/len(moc))
	// students avg for eng
	fmt.Printf("stuends English avg :: %d \n", engtotal/len(moc))

}

// 第一题未完成，可以继续思考下
// 第二题的if可以实现下
// 第三题未实现统计不及格的人数
