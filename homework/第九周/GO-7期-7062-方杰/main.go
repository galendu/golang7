package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

// 遍历目录
func walkDir(path string) ([]string, error) {
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
				if s, err := walkDir(filepath.Join(path, fileInfo.Name())); err != nil {
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

func writeFile(file, mergeFile string) error {
	fmt.Printf("Writing file %s\n", file)
	fout, err := os.OpenFile(mergeFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer fout.Close()
	writer := bufio.NewWriter(fout)

	if fin, err := os.Open(file); err != nil {
		return err
	} else {
		defer fin.Close()
		defer writer.Flush()
		writer.WriteString(fmt.Sprintf("# Content From File '%s': >>\n", file))
		reader := bufio.NewReader(fin)
		for {
			if line, err := reader.ReadString('\n'); err != nil {
				if err == io.EOF {
					writer.WriteString(line)
					writer.WriteString("\n")
					break
				}
				return err
			} else {
				writer.WriteString(line)
			}
		}
	}
	return nil
}

func mergeFiles(files []string, target string) {
	if _, err := os.OpenFile(target, os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(len(files))

	for _, fin := range files {
		go func(fin string) {
			err := writeFile(fin, target)
			if err != nil {
				fmt.Println(err)
			} else {
				wg.Done()
			}
		}(fin)
	}
	wg.Wait()
}

func main() {
	dir, _ := os.Getwd()
	//fmt.Println(dir)
	dataDir := path.Join(dir, "/homework/第九周/GO-7期-7062-方杰/data/")
	target := path.Join(dir, "/homework/第九周/GO-7期-7062-方杰/target.txt")
	if files, err := walkDir(dataDir); err != nil {
		fmt.Println(err)
	} else {
		mergeFiles(files, target)
	}
}
