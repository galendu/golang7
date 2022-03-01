package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

var Loc_time *time.Location

func init() { Loc_time, _ = time.LoadLocation("Asia/Shanghai") }

func time_Time(t string) {
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	if err != nil {
		fmt.Printf("time.ParseInLocation err", err)
		return
	}
	fmt.Println(t1.Format("200601021504"))
}
func class() {
	today := time.Now()
	n := int(time.Saturday - today.Weekday()) //距下次星期六的时间
	if n == 0 {
		n = 7
	}
	nextsaturday := time.Now().AddDate(0, 0, n)
	fmt.Printf("未来第1次上课时间为%.10v\n", nextsaturday)
	for i := 2; i < 5; i++ {
		nextday := nextsaturday.AddDate(0, 0, 7*i)
		fmt.Printf("未来第%d次上课时间为%.10v\n", i, nextday)
	}
}
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

func main() {
	time_Time("1998-10-01 08:10:00")
	class()
	compress("./dir", "test.zip")
}
