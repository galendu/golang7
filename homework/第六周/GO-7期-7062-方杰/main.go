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

// 作业：
// 1.把字符串 "1998-10-01 08:10:00"解析成time.Time格式，再格式化成字符串"199810010810"
// 2.我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定节假日）
// 3.把一个目录下所有.txt文件合并成一个大的.txt文件，然后对这个大文件进行压缩
// 4.自己实现一个BufferedFileWriter

const defaultBufferSize = 10 // 默认缓冲区大小

// 把字符串 "1998-10-01 08:10:00"解析成time.Time格式，再格式化成字符串"199810010810"
func parse_time() {
	str := "1998-10-01 08:10:00"

	TIME_FMT := "2006-01-02 15:04:05"

	timobj, err := time.Parse(TIME_FMT, str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timobj)
	fmt.Println(timobj.Format("200601021504"))

	// 考虑时区，加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	timeObj, err := time.ParseInLocation(TIME_FMT, str, loc)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(timeObj)
	}

	fmt.Println(timeObj.Format("200601021504"))
}

// 我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定节假日） - 需要手动输入上周六的日期
func saturday(lastSaturday string) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}
	last_date_obj, err := time.ParseInLocation("2006-01-02 15:04:05", lastSaturday, loc)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(last_date_obj)

	next_dates_obj := []time.Time{}
	for i := 0; i < 4; i++ {
		new_date_obj := last_date_obj.Add(time.Hour * 7 * 24)
		fmt.Println(new_date_obj.Format("2006-01-02 15:04:05"))
		next_dates_obj = append(next_dates_obj, new_date_obj)
		last_date_obj = new_date_obj
	}
}

// 我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定节假日）- 根据当前时间自动计算
func class_date() {
	now := time.Now()
	// time.Weekday类型可以做运算，强制转int,会得到偏差数
	offset := int(time.Saturday - now.Weekday())

	if offset > 0 { // 今天还没到指定的日期(周六)，那么上一个周六在上周
		offset = offset - 7 // 当前时间偏差 - 7 就是往上周的偏差天数
	}
	saturday := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	fmt.Printf("上次上课日期是: %v\n", saturday.Format("2006-01-02 15:04:05"))

	fmt.Println("后面四次上课日期(不考虑节假日)是:")
	for i := 1; i <= 4; i++ {
		fmt.Println(saturday.Add(time.Duration(i*7*24) * time.Hour).Format("2006-01-02 15:04:05"))
	}
}

// 遍历目录
func walk_dir(path string) ([]string, error) {
	files_path := []string{}
	if fileInfos, err := ioutil.ReadDir(path); err != nil {
		return nil, err
	} else {
		for _, fileInfo := range fileInfos {
			if !fileInfo.IsDir() {
				fileName := fileInfo.Name()
				//fmt.Printf("%T\n", fileName)
				if strings.HasSuffix(fileName, ".txt") {
					filePath := filepath.Join(path, fileName)
					files_path = append(files_path, filePath)
				}
			}
			if fileInfo.IsDir() {
				if s, err := walk_dir(filepath.Join(path, fileInfo.Name())); err != nil {
					return s, err
				} else {
					for _, ele := range s {
						files_path = append(files_path, ele)
					}
					return files_path, err
				}
			}
		}
		return files_path, err
	}
}

// 把一个目录下所有.txt文件合并成一个大的.txt文件，然后对这个大文件进行压缩 - 实现多个文件内容合并为一个文件
func merge_files(files []string, fileName string) {
	for _, file := range files {
		// 创建要写入的文件句柄
		fout, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		}
		writer := bufio.NewWriter(fout)

		// 读文件的文件句柄
		if fin, err := os.Open(file); err != nil {
			fmt.Println(err)
		} else {
			defer fin.Close()
			defer fout.Close()
			defer writer.Flush()
			reader := bufio.NewReader(fin)
			bs := make([]byte, 1024)
			for {
				if n, err := reader.Read(bs); err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println(err)
				} else {
					writer.Write(bs[:n])
				}
			}
		}
	}
}

// 把一个目录下所有.txt文件合并成一个大的.txt文件，然后对这个大文件进行压缩 - 实现大文件压缩
func compress_file(fileName string) string {
	tarName := fmt.Sprintf("%s.zlib", strings.Split(fileName, ".")[0])
	//fmt.Println(tarName)

	if fin, err := os.Open(fileName); err != nil {
		fmt.Println(err)
	} else {
		defer fin.Close()
		if fout, err := os.OpenFile(tarName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
			fmt.Println(err)
		} else {
			defer fout.Close()
			writer := zlib.NewWriter(fout)
			defer writer.Flush()
			for {
				bs := make([]byte, 1024)
				if n, err := fin.Read(bs); err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println(err)
				} else {
					writer.Write(bs[:n])
				}
			}
		}
	}
	return tarName
}

// 查看压缩文件内容
func un_compress(fileName string) {
	if fin, err := os.Open(fileName); err != nil {
		fmt.Println(err)
	} else {
		defer fin.Close()
		if reader, err := zlib.NewReader(fin); err != nil {
			fmt.Println(err)
		} else {
			defer reader.Close()
			fmt.Println("压缩包中文件的内容:")
			io.Copy(os.Stdout, reader) // 直接通过io流拷贝完成读取内容的输出
		}
	}
}

func main() {
	// 1. 时间解析和格式化
	parse_time()

	// 2.计算下4次上课日期
	// 需要手动指定上一个周六的日期
	//lastSaturday := "2022-02-19 09:00:00"
	//saturday(lastSaturday)

	// 无需手动输入上一个周六，自动根据当前时间获取
	//class_date()

	// 3. 合并文件并压缩文件
	/*
		dir, _ := os.Getwd()
		dataDir := path.Join(dir, "/homework/第六周/GO-7期-7062-方杰/data/")
		newFile := path.Join(dir, "/homework/第六周/GO-7期-7062-方杰/merge.txt")

		if files, err := walk_dir(dataDir); err != nil {
			fmt.Println(err)
		} else {
			merge_files(files, newFile)       // 将多个txt文件合并成一个大的txt文件
			tarName := compress_file(newFile) // 压缩这个大txt文件
			un_compress(tarName)              // 读取压缩后的文件内容
		}
	*/

	// 4.自己实现一个BufferedFileWriter
	/*
		filePath, _ := os.Getwd()
		filePath = path.Join(filePath, "/homework/第六周/GO-7期-7062-方杰/out.txt")
		if file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()
			writer := bufw.NewWriterSize(file, 10)
			writer.WriteString("123456789")
			writer.WriteString("\n")
			writer.WriteString("0abc")
			writer.WriteString("\t")
			writer.WriteString("defghijklm")
			defer writer.Flush() // 防止缓冲区还有内容没有被写入文件
		}
	*/
}
