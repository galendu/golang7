package main


import (
	"io"
	"compress/zlib"
	"strings"
	"os"
	"io/ioutil"
	"fmt"
)
//3. 把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩

// 遍历一个目录下有哪些.txt文件
func GetAllFiles(dirpath string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}
	
	pthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs,dirpath + pthSep+fi.Name() )
			GetAllFiles(dirpath + pthSep + fi.Name())
		}else {
			ok := strings.HasSuffix(fi.Name(), ".txt")
			if ok {
				files = append(files, dirpath+pthSep+fi.Name())
			}
		}
	}
	
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

// 复制
func ReadAndWritefile() {
	file1 := "../question3/a.txt"
	file2 := "../question3/b.txt"
	data, err := ioutil.ReadFile(file1)
	if err != nil {
		fmt.Printf("read file fail %v\n", err)
		return
	}
	err = ioutil.WriteFile(file2, data, 0666)
	if err != nil {
		fmt.Printf("write file fail %v\n", err)
	}
}

// 将file的内容追加写入到big.txt
func appendwrite(file string){
	// 如何将遍历出来的文件名，作为值传到file1
	file1 := file
	data, err := ioutil.ReadFile(file1)
	if err != nil {
		fmt.Printf("read file fail %v\n", err)
		return
	}
	d := string(data)


	f, err := os.OpenFile(`../question3/big.txt`, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer f.Close()
	if err !=nil {
		panic(err)
	}

	f.WriteString(d)
}

func compress() {
	fin, err := os.Open("../question3/big.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	stat, _ := fin.Stat()
	fmt.Printf("压缩前文件的大小为:%dB\n", stat.Size())

	fout, err := os.OpenFile("../question3/big.zlib", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, 1024)
	writer := zlib.NewWriter(fout)
	for {
		n, err := fin.Read(bs)
		if err != nil {
			if err == io.EOF {
				break
			}else {
				fmt.Println(err)
			}
		}else {
			writer.Write(bs[:n])
		}
	}
	writer.Close()
	fout.Close()
	fin.Close()

	fin, err = os.Open("../question3/big.zlib")
	if err != nil {
		fmt.Println(err)
		return
	}
	stat, _ = fin.Stat()
	fmt.Printf("压缩后文件大小为:%dB\n", stat.Size())

	fin.Close()
}

func main() {
	xfiles, _ := GetAllFiles("../question3")
	for _, File := range xfiles {
		fmt.Printf("获取的文件为:[%s]\n", File)
	}
	//ReadAndWritefile()

	appendwrite("../question3/a.txt")
	compress()
	

}