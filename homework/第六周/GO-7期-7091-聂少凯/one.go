package main

import (
	"fmt"
	"time"
)

//1. 把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810
func One() {
	str := "1998-10-01 08:10:00"
	TIME_FMT := "2006-01-02 15:04:05"
	layout := "200601021504"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	dt, _ := time.ParseInLocation(TIME_FMT, str, loc)
	dstr := dt.Format(layout)
	fmt.Printf("dt类型是%T\n值是%v\n格式化后是%v\n",dt,dt,dstr)
	//var a string = time.Now().UTC().Add(8 * time.Hour).Format("2006$01$02")
}
