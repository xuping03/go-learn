package main

import "fmt"

func main() {
	// var b [5]int
	// fmt.Println("len(b)", len(b))

	// var a [10]int
	// a[0] = 1
	// i := 1
	// a[i] = 2

	// for i := 0; i < len(a); i++ {
	// 	a[i] = i + 1
	// }
	// for i, v := range a {
	// 	fmt.Printf("a[%d]=%d\n", i, v)
	// }

	// 数组的定义
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a=", a)

	b := [5]int{6, 7, 8, 9, 10}
	fmt.Println("b=", b)

	c := []int{1, 2, 3, 4, 5}
	fmt.Println("c=", c)

	d := [5]int{1, 2, 3}
	fmt.Println("d=", d)

	// 二维数组
	// 只初始化某一个
	e := [3][4]int{1: {5, 6, 7, 8}}
	fmt.Println("e=", e)

	// 数组的比较，支持 == 或 ！=，比较是不是每一个元素都一样，2个数组比较，数组类型要一样
	e1 := [2]int{1, 2}
	e2 := [2]int{1, 2}
	fmt.Println("e1==e2?", e1 == e2)
	e3 := [2]int{2, 1}
	fmt.Println("e1==e3?", e1 == e3)

}
