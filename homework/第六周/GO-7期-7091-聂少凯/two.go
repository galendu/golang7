package main

import (
	"fmt"
	"time"
)

//2. 我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）
func Two(){
	nowDay := time.Now().Weekday()
	offset := int(time.Saturday - nowDay)
	if offset == 0{
		offset = 7
	}
	firstSaturday := time.Now().UTC().AddDate(0,0,offset)
	for i := 1;i<=4 ;i++{
		formatDate := firstSaturday.Format("2006-01-02")
		fmt.Printf("未来第%d次上课日期是%v\n",i,formatDate)
		firstSaturday = firstSaturday.AddDate(0,0,7)
	}
}