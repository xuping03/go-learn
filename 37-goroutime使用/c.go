package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccc")
	// return
	runtime.Goexit()
	fmt.Println("dddddd")
}

func main() {
	go func() {
		fmt.Println("aaaaaaa")
		test()
		fmt.Println("bbbbb")
	}()

	for {

	}
}
