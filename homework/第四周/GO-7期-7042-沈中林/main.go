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
		fmt.Println("-------", num)
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
type Fish interface {
	Swimming(name string)
}

type Cat interface {
	Jumping(name string)
}

type Animal interface {
	Eating(name string)
	Fish
	Cat
}

type Frog struct {
	City string
}

func (f *Frog) Eating(name string) {
	fmt.Printf("Frog %s is eating in %s \n", name, f.City)
}

func (f *Frog) Swimming(name string) {
	fmt.Printf("Frog %s is swimming in %s \n", name, f.City)
}

func (f *Frog) Jumping(name string) {
	fmt.Printf("Frog %s is jumping in %s \n", name, f.City)
}

func animalMethod(name string, ani Animal) {
	ani.Eating(name)
	ani.Jumping(name)
	ani.Swimming(name)
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
	cpMulIncrFunc := ComputeMulIncr(1, 2.0, 4.0)
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
		City: "bg",
	}
	animalMethod("Tom", Tom)

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

