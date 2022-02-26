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

var Loc_time *time.Location

func init() { Loc_time, _ = time.LoadLocation("Asia/Shanghai") }

//1.把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810
func Fmt_time(time1 string) string {
	Time_fmt1 := "2006-01-02 15:04:05"
	Time_fmt2 := "200601021504"
	if t1, err := time.ParseInLocation(Time_fmt1, time1, Loc_time); err == nil {
		t2 := t1.Format(Time_fmt2)
		return t2
	} else {
		return fmt.Sprintf("错误:%v\n", err) //当时间格式Time_fmt1里面出现格式错误时，IDE不会自动报错，只有运行时才会报错。尽量将err都打出来
	}
}

//2.我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）
func class_day(n int) {
	today := time.Now()
	time_dif := time.Saturday - today.Weekday()
    if time_dif==0{
        time_dif=7
    }
	nestsaturday := time.Now().AddDate(0, 0, int(time_dif))
	for i := 0; i < n; i++ {
		nestday := nestsaturday.AddDate(0, 0, 7*i)
		fmt.Printf("未来第%d次课的上课日期是%.10v,\n", i+1, nestday)
	}
}

//3. 把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩
func compose_filer(path string) (path2 string) { //把一个目录下的所有.txt文件合一个大的.txt文件，返回新的文件的目录
	fileinfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("错误0:\n%v\n", err)
	}
	path2 = "G:/大文件.txt"
	Filetxt, err := os.OpenFile(path2, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("错误1:%v\n", err)
	}
	for _, fileinfo := range fileinfos {
		if fileinfo.IsDir() {
			if compose_filer(filepath.Join(path, fileinfo.Name())); err != nil {
				fmt.Printf("错误2:%v\n", err)
			}
		} else {
			if fin, err := os.Open(filepath.Join(path, fileinfo.Name())); err != nil {
				fmt.Printf("错误3:%v\n", err)
			} else {
				defer fin.Close()
				fin.Seek(0, 0)
				reader := bufio.NewReader(fin)
				for {
					if line, err := reader.ReadString('\n'); err != nil {
						if err == io.EOF {
							if len(line) > 0 {
								defer Filetxt.Close()
								writer := bufio.NewWriter(Filetxt)
								writer.WriteString(line)
								writer.WriteString("\n")
								writer.Flush()
							}
							break
						} else {
							fmt.Printf("错误4:%v\n", err)
						}
					} else {
						line = strings.TrimRight(line, "\n")
						defer Filetxt.Close()
						writer := bufio.NewWriter(Filetxt)
						writer.WriteString(line)
						writer.WriteString("\n")
						writer.Flush()
					}
				}
			}
		}
	}
	return path2

}

func compress_filer(path string) { //根据上面返回的新文件的目录打开文件进行压缩
	fin, err := os.Open(path)
	if err != nil {
		fmt.Printf("错误11:%v\n", err)
		return
	}
	stat2, err := fin.Stat()
	if err != nil {
		fmt.Printf("错误12:\n%v\n", err)
	}
	fmt.Printf("文件压缩前大小:%dB\n", stat2.Size())

	fout, err := os.OpenFile(fin.Name()+".zlib", os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("错误13:\n%v\n", err)
	}
	bs := make([]byte, stat2.Size())
	writer := zlib.NewWriter(fout)
	for {
		n, err := fin.Read(bs)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("错误14:\n%v\n", err)
			}
		} else {
			writer.Write(bs[:n])
		}
	}
	writer.Close()
	fout.Close()
	fin.Close()

	fin, err = os.Open(fout.Name())
	if err != nil {
		fmt.Printf("错误15:\n%v\n", err)
		return
	}
	stat2, _ = fin.Stat()
	fmt.Printf("压缩后文件大小:%dB\n", stat2.Size())

}

//4. 自己实现一个BufferedFileWriter
type BufferedFileWriter struct {
	Buffer4      []byte
	Maxbufferlen int
	Write4       io.Writer
}

func NewBufferedFileWriter(Write4 io.Writer, n int) *BufferedFileWriter { //创造一个构造函数
	return &BufferedFileWriter{
		Buffer4:      make([]byte, n),
		Maxbufferlen: n,
		Write4:       Write4,
	}
}

func (bfWriter *BufferedFileWriter) Writerstring4(arr []byte) {
	bfWriter.Buffer4 = append(bfWriter.Buffer4, arr...)
	if len(bfWriter.Buffer4) >= bfWriter.Maxbufferlen {
		bfWriter.Flush4()
	} else {
		return
	}
}

func (bfWriter *BufferedFileWriter) Flush4() {

}

func (bfWriter *BufferedFileWriter) Close4() {

}

func main() {
	//1
	// time1 := "1998-10-01 08:10:00"
	// fmt.Println(Fmt_time(time1))

	//2
	// class_day(4)

	//3
	// path1 := "G:/go/练习/sl/a"
	// compose_filer(path1)
	// path2 := compose_filer(path1)
	// compress_filer(path2)

	//4（暂未做出）

}

//go run homework/6/6.go
