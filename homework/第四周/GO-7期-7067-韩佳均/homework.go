package main

import (
	"errors"
	"fmt"
)

func work1(num ...float64) (float64, error) {
	c := 1.0
	for _, v := range num {
		c *= v
	} //参数乘积
	if c == 0 {
		return 0, errors.New("除数不能为0")
	}
	return 1 / c, errors.New(" ")
}
func work2_1(args ...float64) float64 {
	n := len(args)
	if n == 1 {
		return args[0]
	}
	return args[n-1] * work2_1(args[:n-1]...)
}
func work2_2(args ...float64) (float64, error) {
	r := work2_1(args...)
	if r == 0 {
		return 0, errors.New("除数不能为0")
	}
	return 1 / r, errors.New(" ")
}

type Fish interface {
	swim() error
}
type Reptiles interface {
	move() error
}
type Frog struct {
	name string
}

func (frog Frog) swim() error {
	fmt.Printf("青蛙在水里游\n")
	return errors.New("123")
}
func (frog Frog) move() error {
	fmt.Printf("青蛙在地里走\n")
	return errors.New("123")
}
func Square(num interface{}) interface{} {
	switch n := num.(type) {
	case float32:
		fmt.Println("num的类型是float32")
		return n * n
	case float64:
		fmt.Println("num的类型是float64")
		return n * n
	case int:
		fmt.Println("num的类型是int")
		return n * n
	case byte:
		fmt.Println("num的类型是byte")
		return n * n
	default:
		fmt.Println("num的类型是其他类型")
		return nil
	}

}
func main() {
	w1, err := work1(5.0, 6.1, 7.2, 8.3)
	fmt.Printf("%f%s\n", w1, err)
	w2, err := work2_2(5.0, 6.1, 7.2, 8.3)
	fmt.Printf("%f%s\n", w2, err) //使用递归来实现函数
	var f1 Fish
	var f2 Reptiles
	var f Frog
	f1 = f
	f2 = f
	f1.swim()
	f2.move()
	Square(5.0)
	Square(5)

	Square(5.)

	Square(5)

	Square("sss")

	Square("s")

}
