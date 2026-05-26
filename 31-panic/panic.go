package main

import "fmt"

func testa1() {
	fmt.Println("hhhhhh")
	// panic("hhhhh")
}
func testb(x int) {
	var a [10]int
	a[x] = 111
}

func main1() {
	testa1()
	testb(11)
}
