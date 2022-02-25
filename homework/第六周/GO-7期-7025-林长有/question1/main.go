package main

import (
	"fmt"
	"time"
)

// 1. 把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810

const Time_fmt = "1998-10-01 08:10:00"

func timedemo() {
	now := time.Now()
	fmt.Println("当前时间：", now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", Time_fmt, loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解析时间类型为：%T\n", timeObj)
	fmt.Println("解析时间", timeObj)

	// 将timeObj格式化为199810010810
	timeFmt := timeObj.Format("200601021504")
	fmt.Println("格式化为字符串199810010810:", timeFmt)
}
func main() {
	timedemo()
}