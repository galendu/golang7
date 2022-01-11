package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func zy1() {
	sli := make([]int, 0, 10)   
	for i := 1; i <= 100; i++ { 
		rand.Seed(time.Now().UnixNano())  
		sli = append(sli, rand.Intn(128)) 
	}
	fmt.Printf("%d\n%v\n", len(sli), sli)

	var m = make(map[int]int) 
	for k, y := range sli {   
		m[y] = k 
	}

	fmt.Printf("%d\n%v\n", len(m), m)
}

func main() {
    fmt.Println("hello,world!")
    zy1()
}