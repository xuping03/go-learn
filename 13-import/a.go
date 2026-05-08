package main

import (
	"fmt"
	"go-learn/13-import/otherfun"
)

func main() {
	test()
	a := otherfun.Add(1, 2)
	b := otherfun.Minus(3, 1)
	fmt.Println(a, b)
}
