package main

import (
	"day4/work/work1"
	"day4/work/work2"
	"day4/work/work3"
	"day4/work/work4"
	"fmt"
	"math/rand"
)

func main() {
	ff := make([]float32, 10)
	//初始化10个伪随机数
	for i := 0; i < len(ff); i++ {
		ff[i] = rand.Float32() * 20
	}
	fmt.Println("-----work1-----")
	work1.Work1(ff...)
	work1.Print(ff)
	fmt.Println("-----work2-----")
	work2.Work2(ff...)

	fmt.Println("-----work3-----")
	var Princ_frog work3.Frog

	Princ_frog.Run()
	Princ_frog.Swim()
	fmt.Println("-----work4-----")

	s := work4.Square(1.1)
	fmt.Println(s)
	s = work4.Square(1.)
	fmt.Println(s)
	s = work4.Square(2)
	fmt.Println(s)
	s = work4.Square("sss")
	fmt.Println(s)
	s = work4.Square('s')
	fmt.Println(s)
}
