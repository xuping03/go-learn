package main

import (
	"fmt"
	"os"
)

func WriteFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer f.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i=%d\n", i)
		f.WriteString(buf)
	}
}

func ReadFile(path string) {
	f, _ := os.Open(path)
	defer f.Close()
	buf := make([]byte, 1024*2)
	n, _ := f.Read(buf)
	fmt.Println("buf=", string(buf[:n]))
}

func main() {
	path := "./demo.txt"
	// WriteFile(path)
	ReadFile(path)
}
