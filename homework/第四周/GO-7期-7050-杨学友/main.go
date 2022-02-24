package main

import (
	"errors"
	"fmt"
)

// 1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error

func muti(nums ...float64) (float64, error) {
	result := 0.0
	for _, ele := range nums {
		if result == 0 {
			result = ele
		} else {
			result = result * ele
		}
	}
	if result != 0 {
		return 1 / result, nil
	}
	return -1, errors.New("除零异常")
}

// 2.上题用递归实现
func muti1(nums ...float64) (float64, error) {

	if len(nums) > 1 {
		r, err := muti1(nums[1:]...)
		if err == nil {
			return 1 / nums[0] * r, nil
		} else {
			return r, err
		}

	} else {
		if len(nums) == 0 || nums[0] == 0 {
			return -1, errors.New("除零异常")
		} else {
			return 1 / nums[0], nil
		}
	}

}

// 3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
type Fish interface {
	move()
}

type Reptile interface {
	move()
}

type Frog struct {
	Name string
}

func (frog Frog) move() {
	fmt.Println(frog.Name + " 游啊游，爬呀爬")
}

// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
func square(num interface{}) interface{} {
	switch v := num.(type) {
	case byte:
		return v * v
	case int:
		return v * v
	case float32:
		return v * v
	case float64:
		return v * v
	default:
		return errors.New("data type not support")
	}
}

func main() {
	sum, err := muti(1, 2, 3)
	if err == nil {
		fmt.Println(sum)
	}
	_, err = muti()
	if err != nil {
		fmt.Println(err.Error())
	}

	sum, err = muti1(1, 2, 3)
	if err == nil {
		fmt.Println(sum)
	}
	_, err = muti1()
	if err != nil {
		fmt.Println(err.Error())
	}

	var fish Fish = Frog{"青蛙王子"}

	fish.move()

	var reptile Reptile = Frog{"癞蛤蟆"}

	reptile.move()

	v := square(8)
	e, ok := v.(error)
	if ok {
		fmt.Println(e.Error())
	} else {
		fmt.Println(v)
	}

}
