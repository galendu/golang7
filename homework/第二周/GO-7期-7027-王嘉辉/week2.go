package main

import (
	"fmt"
	"math/rand"
)

func append111() {
	array := make([]int, 0, 10)
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		array = append(array, rand.Intn(128))
		key := array[i]
		if _, exists := m[key]; exists == false {
			m[key] = 1
		} else {
			m[key] += 1
		}
	}
	fmt.Println(array)
	fmt.Println(m)
	fmt.Println(len(m))
}

func bl222(brr []int) {
	for i := 0; i < len(brr); i++ {
		value := brr[i]
		fmt.Println(string(value))
	}

}

func main() {
	arr := []int{2, 4, 6}
	bl222(arr)
}
