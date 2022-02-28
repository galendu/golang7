package main

import (
	"fmt"
	"time"
)

func keyDay()  {
	now := time.Now()
	sub := 6 - int(now.Weekday())
	interval := sub
	if sub == 0 {
		interval = 7
	}

	firstSatuday := now.Add(24*time.Duration(interval)*time.Hour)
	fmt.Println(firstSatuday.Format("2006-01-02"))

	for i := 0;i<3 ;i++ {
		firstSatuday = firstSatuday.Add(24*7*time.Hour)
		fmt.Println(firstSatuday.Format("2006-01-02"))
	}

}