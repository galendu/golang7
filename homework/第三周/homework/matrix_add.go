package main

//matrixAdd 两个8*5的矩阵逐元素相加
func matrixAdd(A [8][5]int, B [8][5]int) [8][5]int {
	sum := [8][5]int{}

	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			sum[i][j] = A[i][j] + B[i][j]
		}
	}

	return sum
}
