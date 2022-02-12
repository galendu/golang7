// 1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
// 2.上题用递归实现
// 3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte

package main

import (
	"fmt"
)

func floatpro(x float64, y ...float64) float64 {
	pro := x
	for _, v := range y {
		pro = pro * v
	}
	if pro == 0 {
		fmt.Println("error")
	} else {
		shu := 1 / pro
		return shu
	}
	return pro
}

func main() {
	ret := floatpro(12.9, 0, 3.2)
	fmt.Println(ret)
}

// 可以再思考优化下逻辑
// 把递归的方式也想一下
