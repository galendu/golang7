// 1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
// 2.上题用递归实现
// 3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte

package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 1.实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
func test1(f ...float64) (float64, error) {
	// fmt.Println(f, reflect.TypeOf(f))
	var result float64 = 1
	for _, num := range f {
		result *= num
	}
	if result == 0 {
		return 0, errors.New("除数不能为0")
	} else {
		return 1 / result, nil
	}
}

// 2.上题用递归实现
func test2(f ...float64) (float64, error) {
	if len(f) == 0 || f[0] == 0 {
		return 0, errors.New("除0异常")
	}
	var result = f[0]
	var error error
	if len(f) > 1 {
		resNext, errNext := test2(f[1:]...)
		result *= resNext
		error = errNext
	} else {
		result = 1 / result
	}
	return result, error
}

// 3.定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口
type fish interface {
	say(conent string)
}

// 爬行动物
type reptile interface {
	run(distance int)
}

// 两栖动物
type amphibian interface {
	fish
	reptile
}

type frog struct {
	color string
}

func (f *frog) say(content string) {
	fmt.Println("咕嘟嘟oO0~" + content)
}

func (f *frog) run(distance int) {
	fmt.Println("跑了" + strconv.Itoa(distance) + "米")
}

func gua(a amphibian) {
	a.say("hi")
	a.run(3)
}

// 4.实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
func square(num interface{}) interface{} {
	switch v := num.(type) {
	case float32:
		return v * v
	case float64:
		return v * v
	case int:
		return v * v
	case byte:
		n := uint16(v) // todo 这种处理方式是不是有问题...
		return n * n
	default:
		fmt.Printf("unsupport type %T\n", v)
		return nil
	}
}

func main() {
	// 1
	s := []float64{1, 2}
	fmt.Println(test1(s...))
	// 2
	fmt.Println(test2(s...))
	// 3
	fg := frog{"green"}
	gua(&fg)
	// 4
	var f32 float32 = 3
	var f64 float64 = 6
	var i int = 1
	var b byte = 97
	var ui uint = 100
	fmt.Println(square(f32))
	fmt.Println(square(f64))
	fmt.Println(square(i))
	fmt.Println(square(b))
	fmt.Println(square(ui))
}

// 第一题思考下是不是少了一种情况，第二题可以优化下
// 第三题虽然扩展的不是很恰当，但很不错
// 第四题错误可以更明白的提示使用者
