package main

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func TimeTransfer(timeStr string)  {
	TIME_FMT := "2006-01-02 15:04:05"	//时间格式
	//now := time.Now()	//当前时间
	//ts := now.Format(TIME_FMT)
	//fmt.Println(ts)
	//loc,_:= time.LoadLocation("Asia/Shanghai")
	//t,_ := time.ParseInLocation(TIME_FMT,ts,loc)
	//fmt.Println(t)
	if formatTime,err:=time.Parse(TIME_FMT,timeStr);err==nil{
		fmt.Printf("字符串:%s,输出：%s\n",timeStr,formatTime)
		fmt.Println(formatTime.Format("20060102150405"))
	}
}

func LearnDay() {
	//1.确定下次上课日期 2022-02-21 09:00:00
	TIME_FMT := "2006-01-02 15:04:05"	//时间格式
	last_learn_day := "2022-02-19 09:00:00"		//上次上课日期
	//loc,_:= time.LoadLocation("Asia/Shanghai")
	if formatTime,err:=time.Parse(TIME_FMT,last_learn_day);err==nil{
		last_learn_day_unix := formatTime.Unix()-3600*8
		for i:=1;i<=4;i++ {
			day_unix := last_learn_day_unix+int64(i*86400*7)
			day := time.Unix(day_unix,0).Format(TIME_FMT)
			fmt.Printf("第%d次上课日期：%s\n",i+6,day)
		}
	}
}

func GetFilesAndDirs(dirPth string)(files []string, dirs []string, err error) {
	dir,err := ioutil.ReadDir(dirPth);
	if err !=nil {
		return nil,nil,err
	}

	PthSep := string(os.PathListSeparator)

	for _,fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs,dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth+PthSep+fi.Name())
		} else {
			ok := strings.HasPrefix(fi.Name(),".txt")
			if ok {
				files = append(files,dirPth+PthSep+fi.Name())
			}
		}
	}
	return files,dirs,nil
}


func GetAllFiles(dirPth string) (files []string,err error) {
	dir,err := ioutil.ReadDir(dirPth)
	if err !=nil {
		return nil,err
	}
	PthSep := string(os.PathSeparator)

	for _,fi := range dir {
		if fi.IsDir() {
			fileTmp,_ :=GetAllFiles(dirPth+PthSep+fi.Name())
			files = append(files,fileTmp...)
		} else {
			ok := strings.HasSuffix(fi.Name(),".txt")
			if ok {
				files = append(files,dirPth+PthSep+fi.Name())
			}
		}
	}
	return files,nil
}

func ReadAllFile(files []string) (file string,err error){
	content := strings.Builder{}
	for _,fi := range files {
		file, err := os.OpenFile(fi, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("文件打开失败，文件名：%s,err:%s\n", fi,err)
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			str,err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			content.WriteString(str)
		}
	}
	return content.String(),err
}

func WriteFile(content string) {
	filePath := "./../big.txt.zip"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	//write := bufio.NewWriter(file)
	//write.WriteString(content)
	//Flush将缓存的文件真正写入到文件中
	//write.Flush()

	writer := zlib.NewWriter(file)
	defer	writer.Close()
	writer.Write([]byte(content))
	writer.Flush()
}

func main()  {
	//1. 把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810
	//TimeTransfer("1998-10-01 08:10:00")
	//2. 我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）
	//LearnDay()
	//3. 把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩(递归遍历.txt文件夹)
	 var files,_ =GetAllFiles("./../../")
	 var content,_ = ReadAllFile(files)
	 WriteFile(content)
	//4. 自己实现一个BufferedFileWriter

}
