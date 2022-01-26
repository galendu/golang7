package main

import "fmt"

type Student struct {
	Name     string
	YuWen    float32
	ShuXue   float32
	YingYu   float32
	AvgScore float32
}

func (stu *Student) Average() {
	stu.AvgScore = (stu.YuWen + stu.ShuXue + stu.YingYu) / 3
	fmt.Printf("AAAA   %s %f\n", stu.Name, stu.AvgScore)
}

type Class struct {
	Students []*Student
	YuWen    float32
	ShuXue   float32
	YingYu   float32
}

func (cls *Class) YuWenAvg() {
	if len(cls.Students) == 0 {
		return
	}
	var sum float32
	for _, stu := range cls.Students {
		sum += stu.YuWen
	}
	cls.YuWen = sum / float32(len(cls.Students))
}

func (cls *Class) Fail() int {
	n := 0
	for _, stu := range cls.Students {
		if stu.AvgScore < 60 {
			fmt.Printf("%s %f\n", stu.Name, stu.AvgScore)
			n++
		}
	}
	return n
}
