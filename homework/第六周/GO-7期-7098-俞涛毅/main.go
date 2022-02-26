package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "path/filepath"

  // "strings"
  "time"
)

var (
  loc *time.Location
)

const (
  TIME_FMT = "2006-01-02 15:04:05"
  DATA_FMT = "2006-01-02"
)

func init() {
  loc, _ = time.LoadLocation("Asia/Shanghai")
}
func q1() {
  // var origins string
  originStr := "1998-10-01 08:10:00"
  os, _ := time.ParseInLocation(TIME_FMT, originStr, loc)
  osf := os.Format("20060102150405")
  fmt.Printf("格式化后为: %s", osf)

}
func q2(t int) {
  cur_date := time.Now()
  cur_week_num := int(cur_date.Weekday())
  cha := 6 - cur_week_num
  cur_unix := cur_date.Unix()

  for i := 0; i < t; i++ {
    if i == 0 {
      if cha == 0 {
        first_zhouliu := int(cur_unix)
        fmt.Printf("第 1 次   周六日期： %s\n", time.Unix(int64(first_zhouliu), 0).Format(DATA_FMT))
      } else {
        first_zhouliu := cha*86400 + int(cur_unix)
        fmt.Printf("第  1 次   周六日期： %s\n", time.Unix(int64(first_zhouliu), 0).Format(DATA_FMT))
      }
    } else {
      next_zhouliu := (cha+7*i)*86400 + int(cur_unix)
      fmt.Printf("第%3d 次   周六日期： %s\n", i+1, time.Unix(int64(next_zhouliu), 0).Format(DATA_FMT))

    }
  }

}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func walk(dirPth, suffix string) (files []string, err error) {
  files = make([]string, 0, 30)
  //忽略后缀匹配的大小写

  err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
    if err != nil { //忽略错误
      return err
    }

    if fi.IsDir() { // 忽略目录
      return nil
    }

    endStr := fi.Name()[len(fi.Name())-4 : len(fi.Name())]
    if endStr == suffix {
      files = append(files, filename)
    }

    return nil
  })

  return files, err
}

func q3(dirPth, suffix string) (fs []string, err error) {

  newFile := filepath.Join(dirPth, "TxtNew.txt")
  os.Remove(newFile)
  files, err := walk(dirPth, suffix)
  if err != nil {
    return files, err
  }
  // 设置写文件

  var newfiles []string
  if fout, err := os.OpenFile(newFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
    fmt.Printf("open file faied: %v\n", err)
  } else {
    defer fout.Close() //别忘了关闭文件句柄
    writer := bufio.NewWriter(fout)

    //开始循环读取文件列表
    for _, file := range files {
      fd, _ := os.Stat(file)
      fm := fd.Mode()

      //判断是否问文件 文件则读取
      if !fm.IsDir() {
        newfiles = append(newfiles, file)
        //打开文件
        fin, _ := os.Open(file)
        fin.Seek(0, 0) //定位到文件开头
        reader := bufio.NewReader(fin)
        for { //无限循环
          if line, err := reader.ReadString('\n'); err != nil { //指定分隔符
            if err == io.EOF {
              if len(line) > 0 { //如果最后一行没有换行符，则此时最后一行就存在line里
                writer.WriteString(line)
              }
              break //已读到文件末尾
            } else {
              fmt.Printf("read file failed: %v\n", err)
            }
          } else {
            writer.WriteString(line)
          }
        }

        defer fin.Close()

      }
      writer.Flush()
    }

  }

  return newfiles, err

}

func q4(txt string) {

  var buffSize = 4

  //OpenFile()比Open()有更多的参数选项。os.O_WRONLY以只写的方式打开文件，os.O_TRUNC把文件之前的内容先清空掉，os.O_CREATE如果文件不存在则先创建，0666新建文件的权限设置
  if fout, err := os.OpenFile("buffer.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
    fmt.Printf("open file faied: %v\n", err)
  } else {
    defer fout.Close() //别忘了关闭文件句柄

    //写文本文件建议使用
    writer := bufio.NewWriterSize(fout, buffSize)
    initLen := 0
    for {
      if buffSize < len(txt) {
        writer.WriteString(txt[initLen:buffSize])
        initLen += buffSize
        buffSize += buffSize
        writer.Flush() //调用Flush强行把缓冲中的所有内容写入磁盘
        continue
      } else {
        endLen := len(txt)
        writer.WriteString(txt[initLen:endLen])
        writer.Flush()
        break
      }

    }

  }
}

func main() {
  // 1. 把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810
  // q1()

  // 2. 我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）
  // times:=10
  // q2(times)
  // 3. 把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩

  files, _ := q3("D:\\gohomework\\test\\day6hw\\a", ".txt")
  fmt.Println("本次读取文件以下文件:")
  for _, vs := range files {
    fmt.Println(vs)
  }

  // 4. 自己实现一个BufferedFileWriter

  inputStr := "我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）"
  q4(inputStr)
}
