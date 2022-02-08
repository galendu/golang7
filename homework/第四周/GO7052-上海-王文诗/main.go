package main

import (
	"errors"
	"fmt"
)

/*
1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
2.上题用递归实现 // 实现不了
3.定义两个接口鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
*/
// 1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
func product(args ...float64) (float64, error) {
	if len(args) == 0{
		return 0,errors.New("parameter is null")
	} else if args[0] == 0{
		return 0,errors.New("parameter is zero")
	}
	var s1 = args[0]
	for i:=1;i < len(args);i++{
		//return 1 / args[i],nil
		s1 = s1 * args[i]
	}
	return 1/s1,nil
}

type animal interface {
	fish
	reptile
}

// fish 鱼类
type fish interface {
	swim(name string)
}

// reptile 爬行动物
type reptile interface {
	climb(name string)
}

// frog 青蛙
type frog struct {}

// swim 鱼类
func (f frog) swim(name string) {
	fmt.Printf("%s 属于鱼类！\n",name)
}

// climb 爬行
func (f frog) climb(name string) {
	fmt.Printf("%s 属于爬行动物！\n",name)
}

func run(a animal,name string) {
	a.swim(name)
	a.climb(name)
}

func square(num interface{}) interface{}{
	switch v :=num.(type) {
	case int:
		return v * v
	case float32:
		return v * v
	case float64:
		return v * v
	case byte:
		return v * v
	default:
		fmt.Printf("type of input error,%T\n",num)
		return nil
	}
}

func main() {
	// 第一题
	p1,err:= product(10)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(p1)
	// 第三题
	//var f1 frog
	//run(f1,"青蛙")
	// 第四题
	//s1 := square(100)
	//fmt.Println(s1)
	//s1 = square(byte(100))
	//fmt.Println(s1)
	//s1 = square(float32(1.23))
	//fmt.Println(s1)
	//s1 = square(float64(1.23))
	//fmt.Println(s1)
}
