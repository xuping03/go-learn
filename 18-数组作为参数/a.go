package main

import (
	"fmt"
)

// 数组作为函数参数，是值传递

func modify(a [5]int) {
	a[1] = 100
}

func modify2(p *[5]int) {
	(*p)[1] = 100
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}
