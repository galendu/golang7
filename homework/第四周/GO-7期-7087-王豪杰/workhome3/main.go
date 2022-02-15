package workhome3

import "errors"

func Work(slice ...float64) (float64, error) {
	number := 1.0
	if len(slice) == 0 {
		return 0, errors.New("参数为空,请输入合法参数")
	}
	for _, slices := range slice {
		if slices == 0 {
			return 0, errors.New("除数不能为0")
		} else {
			number *= slices
		}
	}

	return 1 / number, nil
}

// slices和slice换一下
// 实现下递归方式