package main

import (
	"fmt"
	"time"
)

func main() {
	format := "2006-01-02 15:04:05"
	var timestr = "1998-10-01 08:10:00"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fotmat_time, _ := time.ParseInLocation(format, timestr, loc)
	fmt.Printf("%T\n", fotmat_time)

	format2 := "200601021504"
	fmt.Println(fotmat_time.Format(format2))
}
