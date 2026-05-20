package main

import "fmt"

func testa(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var a [10]int
	a[x] = 111
}

func main() {
	testa(1)
}
