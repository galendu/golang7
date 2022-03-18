package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {

	filePath := "/Users/sean/go/src/Gocache/edu/week9/homework/"
	fileNamestr := "a.txt,b.txt,c.txt"

	//打开合并后文件对象
	fout, err := os.OpenFile(filePath+"result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Errorf("%s文件打开失败", "result.log")
	}
	defer fout.Close()

	splitarr := strings.Split(fileNamestr, ",")

	wg.Add(len(splitarr))

	for _, filename := range splitarr {
		fmt.Printf("这是子协程在进行读写%s\n", filename)
		go func(filename string) {
			defer wg.Done()

			file, err2 := ioutil.ReadFile(path.Join(filePath, filename))
			if err2 != nil {
				fmt.Errorf("文件打开失败")
			}
			fout.Write(file)
			fout.Write([]byte("\n"))
		}(filename)
	}

	wg.Wait()
}
