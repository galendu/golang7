package main

import (
	"fmt"
)

func main() {
	// A := [8][5]int{}
	// B := [8][5]int{}

	// for i := 0; i < 8; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		A[i][j] = rand.Intn(100)
	// 		B[i][j] = rand.Intn(100)
	// 	}
	// }

	// C := matrixAdd(A, B)

	// for i := 0; i < 8; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Printf("%5d", C[i][j])
	// 	}
	// 	fmt.Println()
	// }

	// month := 5
	// fmt.Println(month, jijie3(month))
	// month = 10
	// fmt.Println(month, jijie3(month))
	// month = 1
	// fmt.Println(month, jijie3(month))
	// month = 12
	// fmt.Println(month, jijie3(month))
	// month = 0
	// fmt.Println(month, jijie3(month))
	// month = 13
	// fmt.Println(month, jijie3(month))

	s1 := &Student{
		Name:   "s1",
		YuWen:  36,
		ShuXue: 47,
		YingYu: 58,
	}
	s2 := &Student{
		Name:   "s2",
		YuWen:  76,
		ShuXue: 87,
		YingYu: 58,
	}
	s3 := &Student{
		Name:   "s3",
		YuWen:  36,
		ShuXue: 47,
		YingYu: 18,
	}
	class := Class{
		Students: []*Student{s1, s2, s3},
	}

	for _, stu := range class.Students { //for range 取到的是值拷贝
		// fmt.Printf("%p  %p\n", &class.Students[i], &stu)
		stu.Average()
	}

	class.YuWenAvg()
	fmt.Printf("不及格的学生有%d位\n", class.Fail())
}
