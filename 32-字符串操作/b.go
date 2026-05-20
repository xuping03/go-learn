package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	// 第二位为数，第三位指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcgohello")

	fmt.Println("slice=", string(slice))

	var str string
	str = strconv.FormatBool(false)
	// 'f'指打印格式，以小数方式，-1指小数点位数（紧缩模式），64以float处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)

	// 整型转字符串，常用
	str = strconv.Itoa(6666)
	fmt.Println("str=", str)
	flag, _ := strconv.ParseBool("true")
	fmt.Println("flag", flag)

	a, _ := strconv.Atoi("567")
	fmt.Println("a=", a)

}
