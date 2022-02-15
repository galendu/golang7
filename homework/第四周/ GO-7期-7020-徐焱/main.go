package main

import (
	"errors"
	"fmt"
)

func prod(args ...float64) (float64, error) {
	p := 1.0
	for _, ele := range args {
		p = p * ele
	}
	if p != 0 {
		return 1.0 / p, nil
	} else {
		return 0, errors.New("divide by zero")
	}
}

func prod2(args ...float64) (float64, error) {
	if len(args) == 0 {
		return 0, errors.New("divide by zero")
	}
	first := args[0]
	if first == 0 {
		return 0, errors.New("divide by zero")
	}
	if len(args) == 1 {
		return 1 / first, nil
	}
	remain := args[1:]
	res, err := prod2(remain...)
	if err != nil {
		return 0, err
	}
	return 1 / first * res, nil
}

type Fish interface {
	Swing()
	Breath()
}

type Crawl interface {
	Craw()
}

type Frog struct{}

func (Frog) Swing() {
	fmt.Println("swing")
}

func (Frog) Breath() {
	fmt.Println("breath")
}

func (Frog) Craw() {
	fmt.Println("craw")
}

func (Frog) Talk() {
	fmt.Println("talk")
}

func square(num interface{}) interface{} {
	switch v := num.(type) {
	case float32:
		return v * v //v已经是float32类型的数据了
	case float64:
		return v * v //v已经是float64类型的数据了
	case int:
		return v * v //v已经是int类型的数据了
	case byte:
		return v * v //v已经是byte类型的数据了
	default:
		fmt.Printf("unsupported data type %T\n", num)
		return nil
	}
}

func main() {
	//第1题
	// fmt.Println(prod())
	// fmt.Println(prod(2))
	// fmt.Println(prod(2, 4.9))
	// fmt.Println(prod(2, 0, 4.9))
	// fmt.Println("##############")

	//第2题
	// fmt.Println(prod2())
	// fmt.Println(prod2(2))
	// fmt.Println(prod2(2, 4.9))
	// fmt.Println(prod2(2, 0, 4.9))

	//第3题
	// var ifc1 Fish
	// var ifc2 Crawl

	// frog := Frog{}
	// ifc1 = frog
	// ifc2 = frog

	// ifc1.Swing()
	// ifc2.Craw()

	// frog.Talk()

	//第4题
	var a int = 4
	var b float32 = 4
	var c float64 = 4
	var d byte = 4
	var e uint64 = 4
	fmt.Println(square(a))
	fmt.Println(square(b))
	fmt.Println(square(c))
	fmt.Println(square(d))
	fmt.Println(square(e))
}
