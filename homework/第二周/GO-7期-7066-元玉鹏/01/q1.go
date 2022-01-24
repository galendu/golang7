// 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素
package q1

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate a map and return the length of map
func uniqueKey(slc []int) int {
	ret := make(map[int]int, len(slc)/2)

	for _, value := range slc {
		ret[value] += 1
	}
	return len(ret)
}

// Generates a random array of the specified length
func makeSlice(length int) []int {
	rand.Seed(time.Now().UnixNano()) //set seed to initialize a pseudorandom number generator.

	slc := make([]int, 0, 10) //cap为题目要求

	for i := 0; i < length; i++ {
		slc = append(slc, rand.Intn(128))
	}
	return slc
}

// The entrance
func Q1(count int) int {
	fmt.Println(uniqueKey(makeSlice(count)))
	return uniqueKey(makeSlice(count))
}

/*
	The default `math/rand` number generator is deterministic, so it will give the same output sequence for the same seed value.
	To avoid this, we use current time - time.Now().UnixNano() as the seed.
*/
