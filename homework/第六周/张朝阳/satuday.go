package main

import (
	"fmt"
	"time"
)

func keDay() {
	now := time.Now()
	sub := 6 - int(now.Weekday())
	interval := sub //下一个周距今天隔了几天
	if sub == 0 {
		interval = 7
	}
	firstSatuday := now.Add(24 * time.Duration(interval) * time.Hour)
	fmt.Println(firstSatuday.Format("2006-01-02"))

	for i := 0; i < 3; i++ {
		firstSatuday = firstSatuday.Add(24 * 7 * time.Hour)
		fmt.Println(firstSatuday.Format("2006-01-02"))
	}
}

/**
sub  interval
0		7
1	1
2	2
...
6	6
*/
