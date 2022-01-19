package homework1

import (
	"fmt"
	"math/rand"
)

//创建矩阵
//数组必须初始化长度，确定后不可变
//这里用切片
func createArray(row int, column int) [][]int {
	//行 row 列 column
	//二维切片初始化
	slice := make([][]int, row)
	/*二维切片初始化分配内存后，
	内部的一维slice是没有分配内存的，
	因此要使用二维切片保存数据还需要对一维slice分配内存。*/
	for row := range slice {
		slice[row] = make([]int, column)
		//赋值
		for column := range slice[row] {
			slice[row][column] = rand.Intn(129)
		}
	}
	return slice
}

//格式化输出矩阵
func printArray(slice [][]int) {
	for row := range slice {
		//赋值
		for _, value := range slice[row] {
			fmt.Printf("%6d", value)
		}
		fmt.Println()
	}
}

//矩阵相加
func ArrayPlus(row int, column int) {

	sliceA := createArray(row, column)
	sliceB := createArray(row, column)
	//打印
	fmt.Println("打印数组A")
	printArray(sliceA)

	fmt.Println("打印数组B")
	printArray(sliceB)

	fmt.Println("两个数组开始相加")
	sliceC := make([][]int, row)
	for row := range sliceC {
		sliceC[row] = make([]int, column)
		//赋值
		for column := range sliceC[row] {
			sliceC[row][column] = sliceA[row][column] + sliceB[row][column]
		}
	}
	printArray(sliceC)
}

/*
	打印数组A
	    33    15    71    59     1
	     6    57    44    72    36
	    70    47    34   113    88
	    26    11    21    37    98
	    15    90   104    18   127
	    43    47   120    54   119
	    53   120    91    15    37
	    76    41   119   125    18
	打印数组B
	    13    18    74    67   113
	    19    94   100   127    89
	    21    73   117    23    40
	    17    72   122   103    43
	     3    30    61    28   106
	    36   105     2    31    34
	    75   104   106   118   103
	    38    73     7   124    52
	两个数组开始相加
	    46    33   145   126   114
	    25   151   144   199   125
	    91   120   151   136   128
	    43    83   143   140   141
	    18   120   165    46   233
	    79   152   122    85   153
	   128   224   197   133   140
	   114   114   126   249    70

*/
