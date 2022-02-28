package main

import (
	"fmt"
	"time"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	ch := make(chan time.Time, 1)
	format := "2006-01-02"
	now := time.Now()
	oneDay := time.Second * 86400
	for {
		if int(now.Weekday()) == 6 {
			if v, err := time.Parse(format, now.Format(format)); err == nil {
				ch <- v
				close(ch)
				break
			}
		} else {
			now = now.Add(oneDay)
			fmt.Println("当前now为:", now.Weekday())
		}
	}

	for firstDay := range ch {
		fmt.Printf("第一次上课时间是:%s\n", firstDay.Format(format))
		fmt.Printf("第二次上课时间是:%s\n", firstDay.Add(oneDay*7).Format(format))
		fmt.Printf("第三次上课时间是:%s\n", firstDay.Add(oneDay*7*2).Format(format))
		fmt.Printf("第四次上课时间是:%s\n", firstDay.Add(oneDay*7*3).Format(format))
	}

}
