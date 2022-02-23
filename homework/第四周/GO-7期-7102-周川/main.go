package main

import (
	"errors"
	"fmt"
)

func one(f float64, arg ...float64) (float64, error) {
	all := f
	for _, ele := range arg {
		all *= ele
	}
	if all == 0 {
		return 0, errors.New("错误，乘积为0！")
	} else {
		return (1 / all), nil
	}
}

func two(f float64, arg ...float64) (float64, error) {
	if len(arg) == 0 && f != 0 {
		return (1 / f), nil
	} else if len(arg) == 0 && f == 0 {
		return 0, errors.New("错误，乘积为0！")
	} else {
		result1, err1 := two(f)
		result2, err2 := two(arg[0], arg[1:]...)
		if err1 == nil && err2 == nil {
			return result1 * result2, nil
		} else {
			return 0, errors.New("错误，乘积为0！")
		}
	}
}

type Fisher interface {
	Swim()
}

type Crawler interface {
	Crawl()
}

type Frog struct {
	name string
}

func (Frog) Swim() {
	fmt.Println("frog can swim")
}

func (Frog) Crawl() {
	fmt.Println("frog can crawl")
}

func three() {
	var fish Fisher
	frog1 := Frog{name: "kema"}
	fish = frog1
	fish.Swim()
	var crawl Crawler
	crawl = frog1
	crawl.Crawl()
}

func square(num interface{}) interface{} {
	switch v := num.(type) {
	case float32:
		return v * v
	case float64:
		return v * v
	case int:
		return v * v
	case byte:
		return v * v
	default:
		return errors.New("wrong type")
	}
}

func main() {
	// fmt.Println(one(1., []float64{}...))
	// fmt.Println(two(1, 2, 5, 2))
	// three()
	var a int = 2
	// fmt.Printf("return=%d\n", square(a))
	fmt.Println(square(a))
}
