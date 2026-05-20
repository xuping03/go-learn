package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))

	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "x")
	fmt.Println("buf=", buf)

	fmt.Println(strings.Index("helloabc", "abc"))

	buf = strings.Repeat("go", 3)
	fmt.Println("buf=", buf)

	buf = "hello@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println("s2=", s2)

	arr := []int{1, 2, 3}
	fmt.Println("arr=", arr)

	buf = strings.Trim("     are you ok>    ", " ")
	fmt.Printf("buf=#%s#\n", buf)

	// 去除首尾空格，生成切片
	s3 := strings.Fields("  are you ok")
	for i, data := range s3 {
		fmt.Println(i, ",", data)
	}
}
