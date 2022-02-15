package work1

import (
	"errors"
	"fmt"
)

func Work1(f ...float32) (float32, error) {
	var res float32 = 1.0
	for _, v := range f {
		res = res * v

	}
	if res == 0.0 {
		return 0.0, errors.New("divded by zero")
	} else {
		return 1 / res, nil
	}

}

func Print(ff []float32) {
	if f, err := Work1(ff...); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
}

// 要求的是fload64。另外可以思考下是不是少了一种逻辑判断
