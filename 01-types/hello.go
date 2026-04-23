package main

import "fmt"

func main() {
	var a int
	fmt.Println("a =", a)
	var b, c int
	fmt.Println("b =", b, ", c =", c)
	a = 10

	var d int = 20
	fmt.Println("a =", a, ", d =", d)

	f := 30
	fmt.Println("f =", f)
	fmt.Printf("f type is %T\n", f)

	var aa int
	aa, _ = 30, 40
	fmt.Println("aa =", aa)

}
