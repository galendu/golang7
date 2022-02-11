package main

import (
	"errors"
	"fmt"
)

//1.实现一个函数，接收若干个float64(用不定长参数),返回这些参数乘积的倒数,除数为0时返回error
func culNum(arr ...float64) (num float64, err error) {
	cul := 1.0
	if len(arr) == 0 {
		return 0.0, errors.New("输入有误")
	}
	for i := 0; i < len(arr); i++ {
		cul *= arr[i]
		if cul == 0 {
			return 0, errors.New("相乘结果不能为0")
		} else {
			num = 1 / cul
			err = nil
		}
	}
	return num, err
}

//2. 递归实现
func culNum2(arr ...float64) (num float64, err error) {
	cul := 1.0
	if len(arr) == 0 {
		return 0.0, errors.New("输入有误")
	} else {
		for i := 0; i < len(arr); i++ {
			cul *= arr[i]
			if cul == 0 {
				num = 0.0
				err = errors.New("输入不能为0")
			} else {
				crr := arr[1:]
				culNum2(crr...)
				num = 1 / cul
				err = nil
			}
		}

	}
	return num, err
}

//3. 定义两个接口:鱼类和爬行动物,再定义一个结构体:青蛙,同时实现上述两个接口
type Fish interface {
	Swim()
}

type Reptilian interface {
	Jump()
}

type Frog struct {
	Name string
}

func (fro Frog) Swim() {
	fmt.Printf("%s会游泳\n", fro.Name)
}

func (fro Frog) Jump() {
	fmt.Printf("%s会跳\n", fro.Name)
}

//4.实现函数func square(num interface{}) interface{}，
//计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte

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
		return "输入有误"
	}
}

func main() {
	//第一题实例
	// arr := []float64{1.0, 2.0, 5.0}
	// arr := []float64{0.0, 2.0, 5.0}
	// arr := []float64{}
	// a, err := culNum(arr...)
	// fmt.Println(a, err)

	// 递归实现
	// brr := []float64{1.0, 2.0, 5.0}
	// brr := []float64{}
	// brr := []float64{0.0, 2.0, 5.0}
	// a, err := culNum2(brr...)
	// fmt.Println(a, err)

	//第三题实例
	fr := Frog{Name: "青蛙王子"}
	fr.Swim()
	fr.Jump()

	//第四题实例
	a := square(5)
	fmt.Printf("v=%d type=%T\n", a, a)
	b := square(22.)
	fmt.Printf("v=%f type=%T\n", b, b)

}
