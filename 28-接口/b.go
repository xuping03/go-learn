package main

import "fmt"

// 空接口

// 相当于js里的any，可以接受任何类型的参数
func aaa(arg ...interface{}) {

}
func main() {
	var i interface{} = 1
	fmt.Println("i=", i)
	i = "abc"
	fmt.Println("i=", i)
}
