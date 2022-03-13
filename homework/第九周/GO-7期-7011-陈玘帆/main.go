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

var filePathSlice []string

func init() {
	filePathSlice = make([]string, 0)
}

func GetAllFIlePath(path string) ([]string, error) {
	if fileInfos, err := ioutil.ReadDir(path); err != nil {
		return []string{}, err
	} else {
		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() {
				filePathSlice, _ = GetAllFIlePath(filepath.Join(path, fileInfo.Name()))
			} else {
				if strings.HasSuffix(fileInfo.Name(), ".txt") {
					filePathSlice = append(filePathSlice, filepath.Join(path, fileInfo.Name()))
				}
			}
		}
	}
	return filePathSlice, nil
}

func WriteFile(path, target string) error {
	fout, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer fout.Close()
	writer := bufio.NewWriter(fout)

	fin, err := os.OpenFile(path, os.O_RDONLY, 0000)
	if err != nil {
		return err
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

func MergeFile(filepath []string, target string) (string, error) {
	length := len(filepath)
	if length == 0 {
		return "", errors.New("empty file path slice")
	}
	wg := sync.WaitGroup{}
	wg.Add(length)
	for _, path := range filepath {
		go func(path string) {
			fmt.Println(path)
			err := WriteFile(path, target)
			if err != nil {
				fmt.Printf("write file error, error msg:%s", err)
			} else {
				wg.Done()
			}
		}(path)
	}
	wg.Wait()
	return target, nil
}

func main() {
	allFilePath, err := GetAllFIlePath("./homework/第九周/GO-7期-7011-陈玘帆/")
	if err != nil {
		fmt.Printf("get file path error")
	}
	mergedFilePath, err := MergeFile(allFilePath, "./homework/第九周/GO-7期-7011-陈玘帆/merged_file.txt")
	if err != nil {
		fmt.Printf("merged file error")
	}
	fmt.Printf("new file path %s", mergedFilePath)

}
