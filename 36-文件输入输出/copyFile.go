package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 3 {
		fmt.Println("输入有误，用法: copyFile <源文件> <目标文件>")
		return
	}

	srcFile := list[1]
	dstFile := list[2]
	if srcFile == dstFile {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	sf, err1 := os.Open(srcFile)
	if err1 != nil {
		fmt.Println("err", err1)
		return
	}
	df, err2 := os.Create(dstFile)
	if err2 != nil {
		fmt.Println("err2", err2)
		return
	}
	defer sf.Close()
	defer df.Close()

	buf := make([]byte, 4*1024)
	for {
		n, err := sf.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err", err)
		}
		df.Write(buf[:n])
	}
}
