package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

var (
	loc *time.Location
)

const (
	TIME_FMT = "2006-01-02 15:04:05"
	Day      = "2006-01-02"
	DATE_FMT = "200601020000"
)

// 1. 把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810
func time_Time() {
	dt := "1998-10-01 08:10:00"
	loc, _ = time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(TIME_FMT, dt, loc)
	fmt.Println(t) //解析成time.Time
	t1 := t.Format(DATE_FMT)
	fmt.Println(t1)
	//再格式化成字符串

}

// 2. 我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）
//1.获取本周周六的时间

//2.本周周六的时间+7+7+7+7
func Sat() {
	var weekmonday string
	now := time.Now()
	test := int(time.Saturday - now.Weekday())
	webkStartData := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, test)
	one := webkStartData.Format(Day)
	two := webkStartData.AddDate(0, 0, 7).Format(Day)
	three := webkStartData.AddDate(0, 0, 14).Format(Day)
	four := webkStartData.AddDate(0, 0, 21).Format(Day)
	fmt.Printf("未来第一次上课%v\n", one)
	fmt.Printf("未来第二次上课%v\n", two)
	fmt.Printf("未来第三次上课%v\n", three)
	fmt.Printf("未来第四次上课%v\n", four)
	fmt.Println(weekmonday)

}

// 3. 把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩
func compress(dir string, zip string) {
	src_file := dir + "/" + "test.txt"
	dest_file := dir + "/" + "test.gz"
	//1.先列出指定目录下面所有的txt文件
	file, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	//2.把所有的txt文件读取出来并并写到一个新的文件中
	for i := range file {
		file := file[i].Name()
		file = dir + "/" + file
		f, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		sum_file, err := os.OpenFile(src_file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		}
		defer sum_file.Close()
		sum_file.Write([]byte(f))
	}
	//3.压缩合并后的大文件
	fin, err := os.Open(src_file)
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建压缩文件
	fout, err := os.OpenFile(dest_file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	bs := make([]byte, 1024)
	writer := gzip.NewWriter(fout)
	for {
		n, err := fin.Read(bs)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
			}
		} else {
			writer.Write(bs[:n])
		}
	}
	writer.Close()
	fout.Close()
	fin.Close()
}

// 4. 自己实现一个BufferedFileWriter

func main() {
	// time_Time()
	// Sat()
	compress("./dir", "test.zip")
}
