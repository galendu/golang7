package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateInteger(cap int) func() []int {
	return func() []int {
		randSlice := make([]int, 0, cap)
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < cap; i++ {
			randSlice = append(randSlice, rand.Intn(1000))
		}
		return randSlice
	}
}

func NewSubSlice(rows, cols int) [][]int {
	baseSlice := generateInteger(rows * cols)()
	subSlice := make([][]int, 0, rows)
	for i := 0; i < rows; i++ {
		subSlice = append(subSlice, baseSlice[i*cols:cols*(i+1)])
	}
	return subSlice
}

func addSlice(x, y [][]int, sum *([8][5]int)) interface{} {

	for row := 0; row < len(x); row++ {

		for col := 0; col < len(x[row]); row++ {
			sum[row][col] = x[row][col] + y[row][col]
		}
	}
	return sum
}

func main() {

	slc1 := NewSubSlice(8, 5)
	fmt.Println(slc1)
	time.Sleep(1 * time.Second)
	slc2 := NewSubSlice(8, 5)
	fmt.Println(slc1)
	sum := [8][5]int{}
	result := addSlice(slc1, slc2, &sum)
	fmt.Printf("%[1]v\n%[1]d\n", result)

}
