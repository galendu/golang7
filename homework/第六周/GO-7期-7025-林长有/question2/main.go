package main

import (
	"fmt"
	"time"
)
// 2. 我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）

/*
1 获取当前时间
2 判断是否是周六
  如果是周六进行输出，并输出具体的日期,继续寻找下一次周六
  如果不是周六，循环获取下一天
3 累计输出四次周六
*/
func week(){
		now := time.Now()
		//fmt.Printf("week day :%s , %d\n", now.Weekday().String(), now.Weekday())
		fmt.Println("今天是:",now)
		if now.Weekday()==6{
			sum := 0
			for i:=0; i < 4 ;i++ {
				now = now.AddDate(0, 0, 7)
				sum += 1
				fmt.Println("这一天需要上课",now)
			}
		}else{
			fmt.Printf("今天%s不用上课\n", now.Weekday().String())
			s := now.Weekday()
			d := 6-s
			w := int(d)
			t := now.AddDate(0, 0, w)
			fmt.Println("这一天需要上课:",t)
			sum := 0
			for i := 0; i<3; i++ {
				t = t.AddDate(0,0, 7) 
				sum += 1
				fmt.Println("这一天需要上课:", t)
			}
		}


		// d := time.Duration(24 * time.Hour)
		// t2 := now.Add(d)

		
		// fmt.Printf("week day :%s , %d\n", t2.Weekday().String(), t2.Weekday())
	
	}

func main() {
	week()
}