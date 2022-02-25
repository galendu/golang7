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

func FormantDateStr(dateStr string) (string, error) {
	TIME_FMT := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/shanghai")
	if tm, err := time.ParseInLocation(TIME_FMT, dateStr, loc); err != nil {
		return "", err
	} else {
		TIME_FMT_2 := "20060102150405"
		return tm.Format(TIME_FMT_2), nil
	}
}

func getLessonTime(firstLessonTime string) ([]string, error) {
	TIME_FORMAT := "2006-01-02"
	loc, _ := time.LoadLocation("Asia/shanghai")
	date_string := make([]string, 0, 4)
	if tm, err := time.ParseInLocation(TIME_FORMAT, firstLessonTime, loc); err != nil {
		return date_string, err
	} else {
		var lt time.Time
		if tm.Weekday() == 6 {
			lt = tm
		} else {
			if tm.Weekday() == 0 {
				lt = tm.AddDate(0, 0, -1)
			} else {
				days := int(6 - tm.Weekday())
				lt = tm.AddDate(0, 0, days)
			}
		}

		for i := 1; i < 5; i++ {
			lt = lt.AddDate(0, 0, 7)
			date_string = append(date_string, lt.Format(TIME_FORMAT))
		}

	}
	return date_string, nil
}

var filePathSlice []string

func GetAllFile(path string) ([]string, error) {
	if fileInfos, err := ioutil.ReadDir(path); err != nil {
		return []string{}, err
	} else {
		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() { //如果是目录，就递归子遍历
				filePathSlice, _ = GetAllFile(filepath.Join(path, fileInfo.Name()))
			} else {
				if strings.HasSuffix(fileInfo.Name(), ".txt") {
					filePathSlice = append(filePathSlice, filepath.Join(path, fileInfo.Name()))
				}
			}
		}
		return filePathSlice, nil
	}
}

func WriteFile(filePath []string) (string, error) {
	new_file_path := "go_service/gathered.txt"
	fout, err := os.OpenFile(new_file_path, os.O_CREATE|os.O_TRUNC|os.O_RDWR|os.O_APPEND,
		0666)
	if err != nil {
		return "", err
	}
	defer fout.Close()
	writer := bufio.NewWriter(fout)

	for _, fp := range filePath {
		fin, err := os.OpenFile(fp, os.O_RDONLY, 0000)
		if err != nil {
			return "", err
		}
		reader := bufio.NewReader(fin)
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
	return new_file_path, nil
}

func CompassFile(filePath string) (string, error) {
	fin, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	stat, _ := fin.Stat()
	fmt.Printf("压缩前文件大小 %dB\n", stat.Size())

	newFilePath := strings.Split(filePath, ".txt")[0] + ".zlib"
	if fout, err := os.OpenFile(newFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, (0666)); err != nil {
		return "", nil
	} else {

		bs := make([]byte, 1024)

		writer := zlib.NewWriter(fout)
		for {
			n, err := fin.Read(bs)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return "", err
				}
			} else {
				writer.Write(bs[:n])
			}
		}
		writer.Close()
		fout.Close()
		fin.Close()

		fin, err := os.Open(newFilePath)
		if err != nil {
			return "", err
		}
		stat, _ := fin.Stat()
		fmt.Printf("压缩后文件大小%dB\n", stat.Size())

	}
	return filePath, nil
}

type BufferedFileWriter struct {
	buffer []byte
	maxLen int
	writer io.Writer
}

func (bufferedFileWriter *BufferedFileWriter) Writer(content []byte) (n int, err error) {
	bufferedFileWriter.buffer = append(bufferedFileWriter.buffer, content...)
	if len(bufferedFileWriter.buffer) >= bufferedFileWriter.maxLen {
		bufferedFileWriter.Flush()
	}
	return
}

func (bufferedFileWriter *BufferedFileWriter) Flush() error {
	_, err := bufferedFileWriter.writer.Write(bufferedFileWriter.buffer)
	bufferedFileWriter.buffer = []byte{}
	return err
}

func (bufferedFileWriter *BufferedFileWriter) Close() error {
	err := bufferedFileWriter.Flush()
	return err
}

func CustomBufferWriter(origin_writer io.Writer) *BufferedFileWriter {
	max_line := 1024
	return &BufferedFileWriter{
		buffer: make([]byte, 0, max_line),
		maxLen: max_line,
		writer: origin_writer,
	}
}

func CustomBuf(filePath string) {
	fin, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	cus_writer := CustomBufferWriter(fin)
	cus_writer.Writer([]byte("测试一下"))
	cus_writer.Writer([]byte("\n"))
	cus_writer.Writer([]byte("测试两下"))
	cus_writer.Writer([]byte("测试三下"))
	cus_writer.Flush()

	cus_writer.Close()
}

func main() {
	// 第一题
	s, _ := FormantDateStr("2022-02-13 01:02:03")
	fmt.Println(s)

	// 第二题
	dateString, _ := getLessonTime("2022-02-27")
	fmt.Println(dateString)

	// 第三题
	filePath, err := GetAllFile("./homework/第六周/GO-7期-7011-陈玘帆/")
	fmt.Println(filePath, err)

	newFilePath, err := WriteFile(filePath)
	fmt.Println(newFilePath, err)

	fp, err := CompassFile(newFilePath)
	fmt.Println(fp, err)

	// 第四题
	test_fp := "./homework/第六周/GO-7期-7011-陈玘帆/test.txt"
	CustomBuf(test_fp)
}
