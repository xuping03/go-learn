package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	n := len(list)
	fmt.Println("The number of args:", n)

	for i := 0; i < n; i++ {
		fmt.Printf("OS Args %d: %s\n", i, list[i])
	}
}
