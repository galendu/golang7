package homework

import "fmt"

func Square(num interface{})interface{}{
	switch v := num.(type) {
	case int:
		return v * v
	case float64:
		return v * v
	case float32:
		return v * v
	case byte:
		return v * v
	default:
		fmt.Print("不支持的数据类型%T",num)
		return nil
	}
}
