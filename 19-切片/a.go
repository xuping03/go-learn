package main

import (
	"fmt"
)

// 数组作为函数参数，是值传递

func main() {
	a := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[0:3:5]
	fmt.Println("b len", len(b))
	fmt.Println("b cap", cap(b))
	fmt.Println(b)
	s1 := make([]int, 5, 10)
	fmt.Println("s1=", s1)
	s1 = append(s1, 11)
	fmt.Println("s1=", s1)

	s3 := a[2:5]
	fmt.Println("s3", s3)
	fmt.Println("s3[1]", s3[1])
	s4 := s3[2:7]
	fmt.Println("s4", s4)

}
