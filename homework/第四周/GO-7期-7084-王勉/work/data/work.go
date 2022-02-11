package main

import (
	"errors"
	"fmt"
)

func question_one(arr ...float64) (a float64, err error){
	var chengji = 1.0
	for i := 0; i < len(arr); i++ {
		chengji *= arr[i]
	}
	if (chengji == 0) {
		a = 0
		err = errors.New("乘积为0不能求倒")
	} else {
		a = 1/chengji
		err = nil
	}
	return a,err
}

func digui(arr ...float64) float64{
	if (len(arr)==0) {
		return 0
	}
	if(len(arr) == 1) {
		return arr[0]
	}
	return arr[len(arr)-1] * digui(arr[:len(arr)-1]...)
}

func question_two(arr ...float64) (a float64,err error) {
	var chengji = 1.0
	//获取乘积数据时用递归方式
	chengji = digui(arr...)
	if (chengji == 0) {
		a = 0
		err = errors.New("乘积为0不能求倒")
	} else {
		a = 1/chengji
		err = nil
	}
	return a,err
}

//question_three
type Fish interface {
	Swimming(string) string
}

type Paxingdongwu interface {
	Paxing(string) string
}

type Frog struct {
	Name string
}

func (f Frog) Swimming(){
	fmt.Printf("%s:can swiming\n",f.Name)
}

func (f Frog) Paxingdongwu() {
	fmt.Printf("%s:can crawl\n",f.Name)
}

//question_four
func square(num interface{}) interface{} {
	switch r := num.(type) {
	case float32:
		return r * r
	case float64:
		return r * r
	case int:
		return r * r
	case byte:
		return r * r
	default:
		return "it is error type"
	}
}

func main() {

//1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error   
if daoshu, err := question_one(1,2,3,4); err !=nil {
	fmt.Printf("错误输出:%s\n",err)
} else {
	fmt.Printf("输入参数乘积倒数为%f\n",daoshu)
}

//2.上题用递归实现
if daoshu, err := question_two(1,2,3,4); err !=nil {
	fmt.Printf("错误输出:%s\n",err)
} else {
	fmt.Printf("输入参数乘积倒数为%f\n",daoshu)
}

//3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
var f Frog = Frog{
	Name:   "me",
}

var fish = f
var paxing = f
fish.Swimming()
paxing.Paxingdongwu()

//4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
var input float32 = 11.11
fmt.Println("a float32:: ", square(input))
}

// 第一二题思考下是不是少了一种情况
