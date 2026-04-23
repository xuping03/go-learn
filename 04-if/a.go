package main

import (
	"fmt"
)

func main() {
	var a = 10
	if a == 10 {
		fmt.Println("a is 10")
	} else {
		fmt.Println("a is not 10")
	}

	if b := 10; b == 10 {
		fmt.Println("b is 10")
	} else {
		fmt.Println("a is not 10")
	}
}
