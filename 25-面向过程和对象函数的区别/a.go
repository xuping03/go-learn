package main

import "fmt"

func add1(a, b int) int {
	return a + b
}

type long int

func (tmp long) Add2(other long) long {
	return tmp + other
}
func main() {
	var result int
	result = add1(1, 2)
	fmt.Println("result=", result)
	var n long = 2
	r := n.Add2(1)
	fmt.Println("result=", r)
}
