package main

import "fmt"

func myFun(a int) {
	fmt.Println("a=", a)
}

func myFun2(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}

	for i, data := range args {
		fmt.Println(i, data)
	}
	myFun3(args...)     //传递全部参数
	myFun3(args[2:]...) //从args[2]开始传递参数，包含本身
	myFun3(args[:2]...) // 传递args[0]-args[2]，不包含args[2]
}

func myFun3(args ...int) {
	fmt.Println(args)
}

// 直接返回
func myFun4(a int) int {

	return a + 1
}

// 传统写法
func myFun5(a int) (result int) {
	return a + 1
}

// 推荐写法
func myFun6(a int) (result int) {
	result = a + 1
	return
}

// 多个返回值
func myFun7() (a int, b int, c int) {
	a, b, c = 11, 22, 33
	return
}

func main() {
	// myFun(1)
	// myFun2(1, 2, 3)
	// b := myFun4(1)
	// fmt.Println("b=", b)
	a, b, c := myFun7()
	fmt.Println(a, b, c)
}
