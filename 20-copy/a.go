package main

import (
	"fmt"
)

// 数组作为函数参数，是值传递

func main() {
	a := []int{1, 2}
	b := []int{6, 6, 6, 6, 6}
	copy(b, a)
	fmt.Println(b)

}
