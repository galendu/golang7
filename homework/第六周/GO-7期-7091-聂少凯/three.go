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
)

func Three(dir string) error  {
	f,err := os.OpenFile("bigfile",os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_WRONLY,os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	wf := bufio.NewWriter(f)

	zlibf,err := os.OpenFile("zibfile.zlib",os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_WRONLY,os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer zlibf.Close()
	wzlibf := zlib.NewWriter(zlibf)

	fileSlice,err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _,file :=range  fileSlice{
		fmt.Println(file.Name())
		if file.IsDir(){
			continue
		}
		if ! strings.HasSuffix(file.Name(),".txt"){
			continue
		}
		inpath := path.Join(dir,file.Name())
		fmt.Println(inpath)
		err = writeToFile(inpath,wf,wzlibf)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
return nil
}


func writeToFile(inpath string,w *bufio.Writer,wzip *zlib.Writer) error {
	f,err := os.Open(inpath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	fr := bufio.NewReader(f)
	for {
		line, err := fr.ReadString('\n')
		fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					w.WriteString(line)
					w.WriteString("\n")
					wzip.Write([]byte(line))
					wzip.Write([]byte{'\n'})
				}
				break
			} else {
				w.WriteString(line)
				wzip.Write([]byte(line))
			}
		}
	}
	w.Flush()
	wzip.Flush()
	return nil
}

