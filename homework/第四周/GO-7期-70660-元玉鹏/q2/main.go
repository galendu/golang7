package main

import (
	"errors"
	"fmt"
)

// 最笨的方法
func inverseRecursive(element ...float64) interface{} {
	if len(element) == 0 {
		return errors.New("除零异常") // 思考下错误是否和条件对应
	}
	if len(element) == 1 {
		return float64(1) / element[0]
	} else {
		return inverseRecursive(element[1:]...).(float64) / element[0]

		/*
			nextValue, ok := inverseRecursive(element[:len(element)-1]...).(float64)
			if ok {
				return nextValue/element[0]
			} else {
				panic("类型断言失败")
			}
		*/
	}
}

// 闭包实现,ret 必须设置为1
func inverseRecursive2(ret float64, element ...float64) interface{} {
	if len(element) == 0 {
		return errors.New("除零异常") // 同上
	}
	if len(element) == 1 {
		return ret / element[0]
	} else {
		ret /= element[0]
		return inverseRecursive2(ret, element[1:]...).(float64)
	}
}

func main() {
	fmt.Println(inverseRecursive([]float64{3.1, 4.3, 5.5, 6.66666, 7.000}...))
	fmt.Println(inverseRecursive2(float64(1), []float64{3.1, 4.3, 5.5, 6.66666, 7.000}...))

}

// 同第一题，并且想下是不是少了一种判断逻辑
