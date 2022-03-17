package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var (
	fileSlice []string
	wg        sync.WaitGroup
)

//获取制定路径下的所有文件
func getFilePath(path string) error {
	if fileInfos, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, fileInfo := range fileInfos {
			filePath := fileInfo.Name()
			if fileInfo.IsDir() {
				getFilePath(filepath.Join(path, fileInfo.Name()))
			} else {
				fileSlice = append(fileSlice, path+"/"+filePath)
			}

		}
	}
	return nil
}

//合并文件
func mergeFiles(newFile string) error {
	fout, err := os.OpenFile(newFile, os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fout.Close()
	//
	wg.Add(len(fileSlice)) //加N
	writer := bufio.NewWriter(fout)
	for _, i := range fileSlice {
		go func(file string) {
			defer wg.Done() //优雅关闭协程
			fileOnly, err := os.OpenFile(file, os.O_RDONLY, 0000)
			if err != nil {
				fmt.Println(err)
				return
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

		}(i)

	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	writer.Flush()
	return nil
}

func main() {
	getFilePath("./homework/file")

	mergeFiles("./homework/merge.txt")
	fmt.Println(runtime.NumGoroutine())
}
