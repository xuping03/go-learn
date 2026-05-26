package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "3.14 567 agsdg 1.23 7. 8.9 1sdljgl 6.66 7.8 "

	reg := regexp.MustCompile(`\d\.\d+`)
	if reg == nil {
		fmt.Println("err")
		return
	}
	re := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("re====", re)
}
