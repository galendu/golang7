package main

import (
	"fmt"
	"math/rand"
)

func sum() {
	const N1 = 8
	const N2 = 4
	A := [N1][N2]float64{}

	for m := 0; m < N1; m++ {
		for n := 0; n < N2; n++ {
			A[m][n] = rand.Float64()
		}

	}
	// fmt.Println(A)

	B := [N1][N2]float64{}

	for m := 0; m < N1; m++ {
		for n := 0; n < N2; n++ {
			B[m][n] = rand.Float64()
		}
	}

	sum := 0.
	for i := 0; i < N1; i++ {
		for j := 0; j < N2; j++ {
			sum += A[i][j] + B[i][j]
			// fmt.Println(sum)
		}
	}
	fmt.Printf("两个矩阵的和为: %f\n", sum)
}

func main() {
	sum()
}
