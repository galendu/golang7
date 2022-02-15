package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func GetReciprocal(other ...float64) (float64, error){
	if len(other) == 0 {
		return -1., errors.New("enmpty slice")
	}
	var product float64
	product = 1.
	for _, ele := range other{
		product *= ele
	}

	if product == 0{
		return -1., errors.New("divide by zero")
	}
	return 1 / product, nil
}


func GetReciprocal2(other ...float64) (float64, error) {
	if len(other) == 0 {
		return -1, errors.New("empty slice")
	}

	firstEle := other[0]
	if firstEle == 0 {
		return -1, errors.New("divide by zero")
	}

	if len(other) == 1 {
		if firstEle == 0 {
			return -1, errors.New("divide by zero")
		} else {
			return firstEle, nil
		}
	}

	reciprocal, err := GetReciprocal2(other[1:]...)
	if err != nil {
		return -1, err
	} else {
		return 1 / (reciprocal * firstEle), nil
	}
}

type Fisher interface {
	swim(distance int) (int, error)
}

type Reptiler interface {
	shin(distance int) (int, error)
}

type Frog struct {
	name string
}

func (frog Frog) swim(distance int)(int, error){
	return distance, nil
}

func (frog Frog) shin(distance int)(int, error){
	return distance, nil
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
		return "invalid type of num"
	}
}


func main(){
	s1 := make([]float64, 0, 10)
	rand.Seed(time.Now().Unix())
	for i:=0; i < 10; i++{
		s1 = append(s1, rand.Float64())
	}
	fmt.Println(s1)

	// 第一题
	if reciprocal, err := GetReciprocal(s1...); err != nil{
		fmt.Println(err.Error())
	}else{
		fmt.Printf("get reciprocal of slice element product: %f", reciprocal)
	}

	// 第二题
	if reciprocal, err := GetReciprocal2(s1...); err != nil{
		fmt.Println(err.Error())
	}else{
		fmt.Printf("get reciprocal of slice element product: %f", reciprocal)
	}

	// 第三题
	flog := Frog{"王子"}
	fmt.Printf("青蛙:%s\n", flog.name)
	var fish Fisher
	fish = flog
	fmt.Println(fish.swim(10))
	var reptile Reptiler
	reptile = flog
	fmt.Println(reptile.shin(20))


	// 第四题
	var d int = 4
	fmt.Println(square(d))
	var c float32 = 4
	fmt.Println(square(c))
	var b float64 = 4
	fmt.Println(square(b))
	var a byte = 255
	fmt.Println(square(a))

}