package main

import (
	"errors"
	"fmt"
)

/*
1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
2.上题用递归实现
3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，
interface{}允许是4种类型：float32、float64、int、byte



*/

func homework1(other ...float64) (float64, error) {
	count := 1.0
	for _, v := range other {
		count *= v
	}

	if count == 0 {
		return 1 / count, errors.New("除数是不能为0的,传入的参数中存在非法值0,请注意")
	}

	return 1 / count, errors.New("") //如果数据正常则返回正常数据，错误为空

}

//递归没有实现上述的功能，写的以下代码感觉不太对
func homework2(n float64) (float64, error) {
	//使用递归方法实现第一个作业
	count := 1.0

	if n == 0 {
		return 0, errors.New("参数不能存在0的情况")
	}
	count *= n

	res, err := homework2(n)
	if err != nil {
		fmt.Println(err)
	}

	return 1 / res, errors.New("")
}

//作业3：定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
//定义一个鱼类的接口
type Fisher interface {
	swimming() //游泳行为
}

//定义一个爬行类的接口
type reptiler interface {
	crawl() //爬行行为
}

//定义一个青蛙的结构体并实现上述的两个接口

type frog struct {
	Name string
}

//青蛙实现鱼类的接口
func (f frog) swimming() {
	fmt.Println("小青蛙会游泳!!!!")
}

//青蛙实现爬行类的接口
func (f frog) crawl() {
	fmt.Println("小青蛙会爬行!!!!")
}

//
func homework3() {
	var f frog
	f.swimming()
	f.crawl()
}

//实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，
//interface{}允许是4种类型：float32、float64、int、byte
func square(num interface{}) interface{} {

	var res interface{}
	switch i := num.(type) {
	case float32:
		fmt.Printf("i的类型是%T\n", i)
		fmt.Printf("num的类型是%T\n", num)
		res = float32(i * i)
	case float64:
		fmt.Printf("i的类型是%T\n", i)
		fmt.Printf("num的类型是%T\n", num)
		res = float64(i * i)
	case int:
		fmt.Printf("i的类型是%T\n", i)
		fmt.Printf("num的类型是%T\n", num)
		res = int(i * i)
	case byte:
		fmt.Printf("i的类型是%T\n", i)
		fmt.Printf("num的类型是%T\n", num)
		fmt.Printf("num的类型是%T\n", num)
		res = byte(i * i)
	default:
		fmt.Println("不在允许的计算的四种类型之内")

	}
	return res
}

func main() {
	//作业1：
	h1 := []float64{1.0, 2.0, 3.0, 1.0}
	res1, error := homework1(h1...)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("作业一的结果:%.2f\n", res1)
	}
	//作业2：
	res2, err := homework2(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("res2:", res2)
	//作业3：
	homework3()
	//作业4：
	res3 := square(6)
	fmt.Printf("res3的类型%T,结果是%v\n", res3, res3)
	res4 := square(3.0)
	fmt.Printf("res4的类型%T,结果是%.2f\n", res4, res4)
	res5 := square(byte('d'))
	fmt.Printf("res5的类型%T,结果是%v\n", res5, res5)
}

