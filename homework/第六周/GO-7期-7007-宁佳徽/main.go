package main

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//1. 把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810
func timeFormat() {
	time_str := "1998-10-01 08:10:00"
	layout := "2006-01-02 15:04:05"

	//先把time_str转换成time.Time对象 才能进行Format转换
	if t, err := time.Parse(layout, time_str); err != nil {
		fmt.Println("Parse err:", err)
	} else {
		fmt.Println(t.Format("200601021504"))
	}
}

//2. 我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）
func getClassTime() {
	now := time.Now()

	//由于Weekday是用iota声明的变量，所以本质上是int类型，可用于数字计算,用6减去今天的日期数字
	offset := int(time.Saturday - now.Weekday())

	saturday := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	fmt.Printf("上一次的上课时间为：%v\n", saturday)

	fmt.Println("后面四次上课时间为：")
	for i := 0; i < 4; i++ {
		saturday = saturday.AddDate(0, 0, 7)
		fmt.Println(saturday)
	}

}

//3. 把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩

//申明一个存放文件路径的slice
var fileSlice []string

//获取指定目录下所有以.txt结尾的文件路径
func getFilePath(path string) error {
	if fileInfos, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, fileInfo := range fileInfos {
			//fmt.Println(fileInfo.Name())
			filePath := fileInfo.Name()
			if fileInfo.IsDir() {
				getFilePath(filepath.Join(path, fileInfo.Name()))
			} else if strings.HasSuffix(filePath, ".txt") { //匹配.txt为结尾的文件
				fileSlice = append(fileSlice, path+"/"+filePath)
			}

		}
	}
	return nil
}

func mergeFiles(newFile string) error {
	fout, err := os.OpenFile(newFile, os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fout.Close()

	writer := bufio.NewWriter(fout)

	for _, i := range fileSlice {
		fileOnly, err := os.OpenFile(i, os.O_RDONLY, 0000)
		if err != nil {
			return err
		}
		reader := bufio.NewReader(fileOnly)
		for {
			if line, err := reader.ReadString('\n'); err != nil {
				if err == io.EOF {
					if len(line) > 0 {
						writer.WriteString(line)
						writer.WriteString("\n")
					}
					break
				} else {
					fmt.Printf("read file failed: %v\n", err)
				}
			} else {
				line = strings.TrimRight(line, "\n")
				writer.WriteString(line)
				writer.WriteString("\n")
			}
		}
	}
	writer.Flush()
	return nil
}

func compressedFile(fileName string) {
	fmt.Println(fileName)
	tarName := strings.Split(fileName, ".txt")[0] + ".zlib"
	fmt.Println(tarName)
	fin, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fout, err := os.OpenFile(tarName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	bs := make([]byte, 1024)
	writer := zlib.NewWriter(fout)
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

//第四题，不会

func main() {

	//第一题
	timeFormat()

	//第二题
	getClassTime()

	//第三题
	getFilePath("./homework/file")
	mergeFiles("./homework/merge.txt")
	compressedFile("./homework/merge.txt")

}
