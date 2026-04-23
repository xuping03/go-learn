package main

import (
	"fmt"
)

func main() {
	var ch byte
	ch = 'a'
	fmt.Println("ch", ch)

	//字符串
	var str = "hhhh"
	fmt.Println("str", str, len(str))

	var t complex128
	t = 2.1 + 1.28i
	fmt.Println("t===", t)
	fmt.Println("shi=", real(t), "xu=", imag(t))

	type bigint int64
	var a bigint
	fmt.Printf("aaa=%T\n", a)
	var b int
	fmt.Printf("aaa=%T\n", b)

	c := 10
	// c := 20
	fmt.Println("c==", c)
}
