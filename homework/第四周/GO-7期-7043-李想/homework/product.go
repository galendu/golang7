package homework

import "errors"

func Productrecip(i ...float64)(res float64,err error){
	var p float64 = 1
	for _,j := range i{
		if j == 0{
			return -1,errors.New("divide by zero")
		}else{
			p = j * p
		}
	}
	return p,nil
}