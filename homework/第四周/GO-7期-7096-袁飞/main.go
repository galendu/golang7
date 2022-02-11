package main

import (
	"errors"
	"fmt"
)

/*
	作业:
		1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
		2.上题用递归实现
		3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
		4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
*/

// 1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
func Product(n ...float64) (result float64, err error) {
	result = 1
	for i := 0; i < len(n); i++ {
		if n[i] == float64(0) {
			err = errors.New("参数中不能有0")
			return
		}
		result *= n[i]
	}
	result = 1 / result
	return
}

// 2.上题用递归实现
func Product2(n ...float64) (result float64, err error) {
	// 未实现?
	if len(n) != 0 && n[0] == float64(0) {
		err = errors.New("参数中不能有0")
		return
	}

	if len(n) == 0 && result != 0 {
		result = 1 / result
		return
	}

	fmt.Println(n[0])
	result *= n[0]
	Product2(n[1:]...)
	return
}

// 3. 定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
type Animal interface {
	Fish()
	Reptilian()
}

type Frog struct {
	Name string
}

func (f *Frog) Fish() {
	fmt.Printf("%s 会鱼类的游泳\n", f.Name)
}

func (f *Frog) Reptilian() {
	fmt.Printf("%s 和爬行动物一样能够在陆地上生活\n", f.Name)
}

func (f *Frog) eat() {
	fmt.Printf("%s 吃害虫\n", f.Name)
}

// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
func square(num interface{}) interface{} {

	switch n := num.(type) {
	case float32:
		return n * n
	case float64:
		return n * n
	case int:
		return n * n
	case byte:
		return n * n
	default:
		return "类型错误，无法计算"
	}
}

func main() {
	// 1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
	result, err := Product(16, 24, 3.1415926, 31)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	// // 2.上题用递归实现，未实现
	// result2, err := Product2(16, 24, 3.1415926, 31)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result2)

	// 3. 定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
	var animal Animal
	var frog = Frog{
		Name: "青蛙",
	}
	animal = &frog
	animal.Fish()
	animal.Reptilian()
	frog.eat()

	// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
	n := square(102)
	fmt.Println(n)
}
