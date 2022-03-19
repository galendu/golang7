package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"sync"
)

//把多个文件合成一个文件，并发实现，不用考虑内容的顺序，子协程要优雅退出
var filePathSlice []string

func init() {
	filePathSlice = make([]string, 0)
}

//1.列出要合并的目录下面所有的文件
func GetFile(path string) ([]string, error) {

	if File, err := ioutil.ReadDir(path); err != nil {
		return []string{}, err
	} else {
		for _, file := range File {
			if file.IsDir() {
				filePathSlice, _ = GetFile(filepath.Join(path, file.Name()))

			} else {
				filePathSlice = append(filePathSlice, filepath.Join(path, file.Name()))
			}

		}
	}
	return filePathSlice, nil
}

//2.写文件
func WriterFile(src, dest string) error {
	fout, err := os.OpenFile(dest, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer fout.Close()
	writer := bufio.NewWriter(fout)

	fin, err := os.OpenFile(src, os.O_RDONLY, 0000)
	if err != nil {
		return nil
	}
	defer fin.Close()
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
				return err
			}
		} else {
			line = strings.TrimRight(line, "\n")
			writer.WriteString(line)
			writer.WriteString("\n")

		}
	}
	writer.Flush()
	return nil
}

//2.把所有的文件写到一个大文件里面
func MergeFile(src []string, dest string) (string, error) {
	length := len(src)
	if length == 0 {
		return "", errors.New("empty file path slice")
	}
	wg := sync.WaitGroup{}
	wg.Add(length)
	for _, file := range src {
		go func(file string) {
			fmt.Printf("正在把%s的文件写入到%s中", file, dest)
			err := WriterFile(file, dest)
			if err != nil {
				fmt.Printf("write file error msg:%s", err)
			} else {
				wg.Done()
			}
		}(file)
	}
	wg.Wait()
	return dest, nil

}

func main() {
	FilePath, err := GetFile("./")
	if err != nil {
		fmt.Printf("get file path error %s\n", err)
	}
	merfilepath, err := MergeFile(FilePath, "./test/merge.txt")
	if err != nil {
		fmt.Printf("merged file error,%s\n", err)
	}
	fmt.Printf("new file pat %s\n", merfilepath)
}
