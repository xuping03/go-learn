package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("a=%d\n", a)
	var p *int
	p = &a
	fmt.Println("p", p)
	fmt.Println("*p", *p)
	*p = 666
	fmt.Println("a", a)

	var cc *int
	cc = new(int)
	*cc = 999
	fmt.Println("*cc", *cc)

}
