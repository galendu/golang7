package main

import (
	"errors"
	"fmt"
)

func prod(a ...float64) (float64, error) {
	p := 1.0
	for _, ele := range a {
		p = p * ele
	}
	if p == 0 {
		return 0, errors.New("devide by zero")
	} else {
		return 1.0 / p, nil
	}
}

func prod2(a ...float64) (float64, error) {
	if len(a) == 0 {
		return 0, errors.New("empty slice")
	}
	first := a[0]
	if first == 0 {
		return 0, errors.New("first num is zero")
	}
	if len(a) == 1 {
		return 1 / first, nil
	}
	a = a[1:]
	result, err := prod2(a...)
	if err != nil {
		return 0, err
	}
	return result / first, nil
}

type Fish interface {
	Swim()
	Brathe()
}

type Crawl interface {
	Craw()
}

type Frog struct {
}

func (Frog) Swim()   {}
func (Frog) Brathe() {}
func (Frog) Craw()   {}

func square(num interface{}) interface{} {
	switch a := num.(type) {
	case float32:
		return a * a
	case float64:
		return a * a
	case int:
		return a * a
	case byte:
		return a * a
	default:
		fmt.Printf("unsupported data type %T\n", num)
	}
	return nil
}

func main() {
	// fmt.Println(prod(2))
	// fmt.Println(prod(2, 3))
	// fmt.Println(prod(2, 5, 0))
	// fmt.Println(prod(2, 5, 10))

	// fmt.Println(prod2(2))
	// fmt.Println(prod2(2, 3))
	// fmt.Println(prod2(2, 5, 0))
	// fmt.Println(prod2(2, 5, 10))

	var a Fish
	var b Crawl
	frog := Frog{}
	a = frog
	b = frog
	a.Swim()
	b.Craw()
	s := square("33")
	fmt.Println(s)
}
