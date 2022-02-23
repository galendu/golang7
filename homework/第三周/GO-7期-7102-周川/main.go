package main

import (
	"fmt"
	"math/rand"
)

func one() {
	a := [8][5]int{}
	b := [8][5]int{}
	for row, s := range a {
		for col := range s {
			a[row][col] = rand.Intn(100)
		}
	}
	for row, s := range b {
		for col := range s {
			b[row][col] = rand.Intn(100)
		}
	}
	sum := 0
	for row, s := range a {
		for col := range s {
			sum = sum + a[row][col] + b[row][col]
		}
	}
	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)
	fmt.Printf("sum=%d\n", sum)
}

func tone(i int) {
	switch i {
	case 2, 3, 4:
		fmt.Println("春天")
	case 5, 6, 7:
		fmt.Println("夏天")
	case 8, 9, 10:
		fmt.Println("秋天")
	case 11, 12, 1:
		fmt.Println("冬天")
	default:
		fmt.Println("请输入1-12")
	}
}

func ttwo(i int) {
	if i >= 2 && i <= 4 {
		fmt.Println("春天")
	} else if i >= 5 && i <= 7 {
		fmt.Println("夏天")
	} else if i >= 8 && i <= 10 {
		fmt.Println("秋天")
	} else if i == 1 || i == 11 || i == 12 {
		fmt.Println("冬天")
	} else {
		fmt.Println("请输入1-12")
	}
}

type student struct {
	name    string
	math    int
	chinese int
	english int
}

func (s student) ave() float64 {
	v := float64((s.math + s.chinese + s.english)) / 3
	return v
}

func (s *student) new(name string) {
	s.name = name
	s.math = rand.Intn(100)
	s.chinese = rand.Intn(100)
	s.english = rand.Intn(100)
}

func (s *student) sum() int {
	return s.math + s.chinese + s.english
}

func three() {
	var u1, u2, u3, u4, u5, u6, u7, u8, u9, u10 student
	class := []student{u1, u2, u3, u4, u5, u6, u7, u8, u9, u10}
	usli := []string{"u1", "u2", "u3", "u4", "u5", "u6", "u7", "u8", "u9", "u10"}
	for i := 0; i < len(class); i++ {
		class[i].new(usli[i])
	}
	var allave float64
	sumall := 0
	bjg := make([]student, 0, 10)
	for _, ele := range class {
		sumall += ele.sum()
		if ele.ave() < 60 {
			bjg = append(bjg, ele)
		}
	}
	allave = float64(sumall) / float64(len(class)*3)
	numbjg := len(bjg)
	fmt.Printf("class=%v\n", class)
	fmt.Printf("allave = %.2f\n", allave)
	fmt.Printf("big=%v\n", bjg)
	fmt.Printf("numbjg=%d\n", numbjg)

}

func main() {
	// one()
	// tone(5)
	// ttwo(0)
	// u1 := student{name: "one", math: 90, chinese: 90, english: 100}
	three()
}
