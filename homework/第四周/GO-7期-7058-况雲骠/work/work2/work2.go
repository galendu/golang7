package work2

import "errors"

func Work2(args ...float32) (float32, error) {
	if len(args) == 0 {
		return 0, errors.New("divide by zero") // 这种逻辑思考下错误是不是对应的条件
	}
	first := args[0]
	if first == 0 {
		return 0, errors.New("divide by zero")
	}
	if len(args) == 1 {
		return 1 / first, nil
	}
	remain := args[1:]
	res, err := Work2(remain...)
	if err != nil {
		return 0, err
	} else {
		return 1 / first * res, nil
	}
}

// 逻辑可以再优化下
