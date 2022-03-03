package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func work1() {
	format1 := "2006-01-02 15:04:05"
	format2 := "20060101150405"
	s1 := "1998-10-01 08:10:00"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(format1, s1, loc)
	s2 := t.Format(format2)
	fmt.Println(s2)
}
func work2() {
	now := time.Now()
	sub := 6 - now.Weekday()
	if sub == 0 {
		sub = 7
	}
	firstsatuday := now.Add(24 * time.Duration(sub) * time.Hour)
	fmt.Println(firstsatuday.Format("2006-01-02"))

	for i := 0; i < 3; i++ {
		firstsatuday = firstsatuday.Add(24 * 7 * time.Hour)
		fmt.Println(firstsatuday.Format("2006-01-02"))
	}
}
func work3(contr11 string) {
	fout, err := os.OpenFile("big.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()

	writer := bufio.NewWriter(fout)
	if files, err := ioutil.ReadDir(contr11); err != nil {
		fmt.Println(err)
	} else {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			baseName := file.Name()
			if strings.HasSuffix(baseName, ".txt") {
				if fin, err := os.Open(filepath.Join(contr11, baseName)); err != nil {
					fmt.Println(err)
					continue
				} else {
					defer fin.Close()
					reader := bufio.NewReader(fin)
					for {
						if line, err := reader.ReadString('\n'); err != nil {
							if err == io.EOF {
								if len(line) > 0 {
									writer.WriteString(line)
									writer.WriteString("\n")
								}
							}
							break
						} else {
							writer.WriteString(line)
						}
					}
				}
			}

		}
	}
	writer.Flush()
}

type Buffered struct {
	buffer      [1024]byte
	fileHandler *os.File
	endpos      int
}

func newbuffered(fd *os.File) *Buffered {
	return &Buffered{
		fileHandler: fd,
	}
}
func (writer *Buffered) Flush() {
	if writer.endpos > 0 {
		writer.fileHandler.Write(writer.buffer[0:writer.endpos])
		writer.endpos = 0
	}
}
func (writer *Buffered) write(content []byte) {
	if len(content) >= 1024 {
		writer.fileHandler.Write(content)
	} else {
		if writer.endpos+len(content) >= 1024 {
			writer.Flush()
			writer.write(content)
		} else {
			copy(writer.buffer[writer.endpos:], content)
			writer.endpos = len(content)
		}
	}
}

func (writer *Buffered) WriteString(content string) {
	writer.write([]byte(content))
}

func testbufferwriter() {
	fout, err := os.OpenFile("123.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()
	writer := newbuffered(fout)
	for i := 0; i < 500; i++ {
		writer.WriteString("0123456789\n")
	}
}
func main() {
	// work1()
	// work2()
	// work3("/dir")
	testbufferwriter()
}
