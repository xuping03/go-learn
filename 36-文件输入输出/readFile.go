package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile(path string) {
	f, _ := os.Open(path)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		fmt.Println(string(buf))
	}

}
func main() {
	path := "./demo.txt"
	readFile(path)
}
