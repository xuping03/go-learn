package main

import "fmt"

func aa(a int) {
	result := 100 / a
	fmt.Println("result", result)
}

func main() {
	defer fmt.Println("bbb")
	defer fmt.Println("aaa")
	defer aa(0)
	defer fmt.Println("ccc")
}
