// 1. 把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810
// 2. 我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）
// 3. 把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩
// 4. 自己实现一个BufferedFileWriter

package main

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)
const TIME_FMT string = "2006-01-02 03:04:05"
const TIME_FMT_WEEK string = "2006-01-02 03:04:05 Monday"

func test1() {
	str := "1998-10-01 08:10:00"
	t, err := time.Parse(TIME_FMT, str)
	if err != nil {
		return
	}
	fmt.Println(t.Format("200601020304"))
}

func test2(n int) {
	today := time.Now()
	year, month, day := today.Date()
	delta := time.Saturday - time.Now().Weekday()
	nextSaturday := time.Date(year, month, day + int(delta), 9, 0, 0, 0, time.Local)
	for i := 0; i < n; i++ {
		nextDay := nextSaturday.AddDate(0, 0, 7 * i)
		fmt.Println(nextDay.Format(TIME_FMT_WEEK))
	}
}

func findFileWithExt(sourceDir, ext string) []string {
	var result []string
	// 1. 遍历目录
	fileInfos, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		return result
	} else {
		for _, fileinfo := range fileInfos {
			name := fileinfo.Name()
			if path.Ext(name) == ext {
				fullname := path.Join(sourceDir, name)
				result = append(result, fullname)
			} else if fileinfo.IsDir() {
				nextDir := path.Join(sourceDir, fileinfo.Name())
				result = append(result, findFileWithExt(nextDir, ext)...)
			}
		}
		return result
	}
}

func readAllFiles(filepaths []string) string {
	sb := strings.Builder{}
	for _, filepath := range filepaths {
		f, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		reader := bufio.NewReader(f)
		for { //无限循环
			if line, err := reader.ReadString('\n'); err != nil { //指定分隔符
				if err == io.EOF {
					if len(line) > 0 { //如果最后一行没有换行符，则此时最后一行就存在line里
						// fmt.Println(line)
						sb.WriteString(line + "\n")
					}
					break //已读到文件末尾
				} else {
					fmt.Printf("read file failed: %v\n", err)
				}
			} else {
				sb.WriteString(line)
			}
		}
	}
	return sb.String()
}

func writeFile(content, targetFilePath string) error {
	file, err := os.OpenFile(targetFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// writer := bufio.NewWriter(file)
	// writer.WriteString(content)
	// writer.Flush()
	writer := zlib.NewWriter(file)
	defer	writer.Close()
	writer.Write([]byte(content))
	writer.Flush()
	return nil
}

func test3(sourceDir string) {
	allTxt := findFileWithExt(sourceDir, ".txt")
	content := readAllFiles(allTxt)
	target := path.Join(sourceDir, "../target/target.txt.gz")
	err := writeFile(content, target)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("write done")
}

func Read(gzPath string) {
	gzfile, err := os.OpenFile(gzPath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("error:", err)
	}
	reader, err := zlib.NewReader(gzfile)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)
}

type BufferFileWriter struct {
	buffer []byte
	maxBufferLen int
	outstream io.Writer
}

func (bfWriter *BufferFileWriter)Write(content []byte) (n int, err error) {
	bfWriter.buffer = append(bfWriter.buffer, content...)
	fmt.Println("---\nwrite - buffer content: ", string(bfWriter.buffer))
	if len(bfWriter.buffer) >= bfWriter.maxBufferLen {
		bfWriter.Flush()
	}
	return
}

func (bfWriter *BufferFileWriter)Flush() error {
	fmt.Println("flush - content: " + string(bfWriter.buffer))
	_, err := bfWriter.outstream.Write(bfWriter.buffer)
	bfWriter.buffer = bfWriter.buffer[:0]
	return err
}

func (bfWriter *BufferFileWriter)Close() error {
	err := bfWriter.Flush()
	fmt.Println("close")
	return err
}

func NewBufferFileWriter(outstream io.Writer) *BufferFileWriter {
	maxBufferLen := 16
	return &BufferFileWriter{
		buffer: make([]byte, 0, maxBufferLen),
		maxBufferLen: maxBufferLen,
		outstream: outstream,
	}
}

func test4(filepath string) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bfw := NewBufferFileWriter(file)
	bfw.Write([]byte("hello"))
	bfw.Write([]byte(" world, 这里凑够16个字"))
	bfw.Write([]byte("   16之后的"))
	defer bfw.Close()
}


func main() {
	// 1
	test1()
	// // 2
	test2(4)
	// 3
	test3("./homework/第六周/GO-7期-7099-佳琦/files")
	// Read()
	// 4
	test4("./homework/第六周/GO-7期-7099-佳琦/target/bfwritecontent.txt")
}