package homework

import (
	"errors"
	"fmt"
)

func Recurproductrecip(arr ...float64)(res float64,err error){
	first := arr[0]
	if first == 0{
		return 1,errors.New("divide by zero")
	}
	if len(arr) == 1{
		return 1/first,nil
	}
	remain := arr[1:]
	fmt.Println(remain)
	result,err := Recurproductrecip(remain...)
	if err != nil{
		return 1,err
	}
	return 1/first * result,nil
}
