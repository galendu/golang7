package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func putinnum(S []int) []int {
	for i := 0; i < 100; i++ {
		S = append(S, rand.Intn(128))
	}
	fmt.Printf("S=%v\n", S)
	return S
}

func mapcreate(M map[int]int, S []int) {
	for j := 0; j < len(S); j++ {
		if value, exists := M[S[j]]; exists {
			M[S[j]] = value + 1
		} else {
			M[S[j]] = 1
		}
	}
	for key, value := range M {
		fmt.Printf("%v=%v\t", key, value)
	}
	fmt.Printf("\n")
	// for _, n := range listOfNumberStrings {
	// 	fmt.Printf("%s\n", *n)
}

func arr2string(arr []int) string {
	mystring := ""
	// for _, ele := range arr {
	// 	//fmt.Printf("element=%d\n", ele)
	// 	mystring = strings.Join([]string{strconv.Itoa(ele)}, " ")
	// }
	for i := 0; i < len(arr); i++ {
		//mystring = strings.Join([]string{strconv.Itoa(arr[i])}, " ")
		mystring += strconv.Itoa(arr[i]) + " "
	}
	return mystring
}

func main() {
	S := make([]int, 0, 10)
	S = putinnum(S)
	fmt.Printf("SE=%v\n", S)
	M := make(map[int]int, 10)
	mapcreate(M, S)
	var inputarr = [...]int{3, 2, 6, 5, 4, 8}
	mystr := arr2string(inputarr[:])
	fmt.Printf("str=%s\n", mystr)

}
