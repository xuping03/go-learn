package main

import (
	"fmt"
)

func main() {
	// 默认有break，需要fallthrough可以继续执行下面的case
	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	for i, data := range str {
		fmt.Println(i, data)
	}
}
