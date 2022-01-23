package main

import (
	"errors"
	"fmt"
)

// 1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
func ComputeMul(num ...float64) (float64, error) {
	cpMul := 1.0
	for _, v := range num {
		cpMul *= v
	}
	if cpMul == 0 {
		return 0.0, errors.New(fmt.Sprintf("paras must > 0, but paras :: %v", num))
	}
	return 1 / cpMul, nil
}

// 2.上题用递归实现
func ComputeMulIncr(num ...float64) func() (float64, error) {
	cpMulIcr := 1.0

	return func() (float64, error) {
		for _, v := range num {
			cpMulIcr *= v
		}
		if cpMulIcr == 0 {
			return 0.0, errors.New(fmt.Sprintf("paras must > 0, but paras :: %v", num))
		}
		return 1 / cpMulIcr, nil
	}
}

// 3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
type Fisher interface {
	Swimming(swimmingPool string)
}

type Cater interface {
	Jumping(place string)
}

type Animaler interface {
	Eating(food string)
	Fisher
	Cater
}

type Frog struct {
	Name string
	City string
}

func (f *Frog) Eating(food string) {
	fmt.Printf("The %s Frog %s is eating food %s \n", f.City, f.Name, food)
}

func (f *Frog) Swimming(place string) {
	fmt.Printf("The %s Frog %s is swimming in %s \n", f.City, f.Name, place)
}

func (f *Frog) Jumping(swimmingPool string) {
	fmt.Printf("The %s Frog %s is jumping in %s \n", f.City, f.Name, swimmingPool)
}

func AnimalMethod(ani Animaler, food, place, swimmingPool string) {
	ani.Eating(food)
	ani.Jumping(place)
	ani.Swimming(swimmingPool)
}

// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
func Square(num interface{}) interface{} {
	switch num.(type) {
	case float32:
		return num.(float32) * num.(float32)
	case float64:
		return num.(float64) * num.(float64)
	case int:
		return num.(int) * num.(int)
	case byte:
		return num.(byte) * num.(byte)
	default:
		return "num type must be float32、float64、int、byte"
	}
}

func main() {
	// 1
	if cpMul, err := ComputeMul(5.0, 2.0, 4.0); err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(cpMul)
	}

	// 2
	cpMulIncrFunc := ComputeMulIncr(1, 0, 4.0)
	if cpMulIncr, err := cpMulIncrFunc(); err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(cpMulIncr)
	}
	if cpMulIncr, err := cpMulIncrFunc(); err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(cpMulIncr)
	}

	// 3
	Tom := &Frog{
		Name: "Tom",
		City: "bg",
	}
	AnimalMethod(Tom, "worm", "jd", "bg")

	// 4
	var a float32 = 12
	fmt.Println("a float32:: ",Square(a))
	var b float64 = 25
	fmt.Println("b float64:: ",Square(b))
	var c int = 100
	fmt.Println("d int:: ",Square(c))
	var d byte = 3
	fmt.Println("d byte:: ",Square(d))

	var e rune = 5
	fmt.Println("e rune:: ",Square(e))
}

