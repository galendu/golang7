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
)

func walk(fileNameSlice *[]string, pathstr string) {
	/*
		获取符合条件的文件名
	*/
	if fileInfos, err := ioutil.ReadDir(pathstr); err != nil {
		os.Exit(1)
	} else {
		for _, fileInfo := range fileInfos {
			if strings.Contains(fileInfo.Name(), ".txt") {
				*fileNameSlice = append(*fileNameSlice, filepath.Join(pathstr, fileInfo.Name()))
			}
			if fileInfo.IsDir() {
				walk(fileNameSlice, filepath.Join(pathstr, fileInfo.Name()))
			}
		}
	}
}

func Gzipfunc(inputName string) (outputName string) {
	/*
		实现压缩 返回压缩后的文件名
	*/
	var sb = strings.Builder{}
	if fin, err := os.OpenFile(inputName, os.O_RDONLY, 0644); err == nil {
		defer fin.Close()
		reader := bufio.NewReader(fin)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Errorf("读取错误 %s", err)
			}
			if err == io.EOF {
				if len(line) > 0 {
					//fmt.Printf("当前读取到文件末尾: %s\n", line)
					sb.WriteString(line)
				}
				break
			} else {
				//fmt.Printf("当前读取行为: %s\n", line)
				sb.WriteString(line)
			}

		}
	}

	fout, err := os.OpenFile(inputName+".zlib", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Errorf("创建、打开待写入文件失败,%s", err)
	}
	writer := zlib.NewWriter(fout)
	//fmt.Printf("待压缩的数据为:%s\n", sb.String())
	writer.Write([]byte(sb.String()))
	fout.Close()
	return inputName + ".zlib"
}

func UnGzipfunc(inputName string) {
	if fin, err := os.Open(inputName); err != nil {
		fmt.Errorf("打开文件失败:%s", err)
	} else {
		defer fin.Close()
		reader, err := zlib.NewReader(fin)
		if err != nil {
			fmt.Errorf("打开带解压文件失败:%s", err)
		}
		io.Copy(os.Stdout, reader)
	}

}

func main() {
	var fileNameSlice = []string{}
	var sb = strings.Builder{}
	pathStr := "edu/week6/homework/Q3/resources"
	walk(&fileNameSlice, pathStr)

	finishedFileName := "result.log"
	finishedFilePath := filepath.Join(pathStr, finishedFileName)
	for index, filename := range fileNameSlice {
		//fmt.Println(filename)
		txtContent, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("读取发生错误:%s\n", err)
		}
		sb.WriteString(string(txtContent))
		if index <= len(fileNameSlice)-1 {
			sb.WriteString("\n")
		}
	}

	err := ioutil.WriteFile(finishedFilePath, []byte(sb.String()), 0644)
	if err != nil {
		fmt.Printf("文件写入发生错误:%s\n", err)
	}

	gipfinishedName := Gzipfunc(finishedFilePath)
	fmt.Println(gipfinishedName)

	UnGzipfunc("gipfinishedName")
}
