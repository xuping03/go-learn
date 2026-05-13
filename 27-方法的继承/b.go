package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Printf("setinfovalue:%p,%v\n", &p, p)
}
func (p *Person) setInfoPointer() {
	fmt.Printf("setinfopointer:%p,%v\n", p, p)
}

func main() {
	p := Person{"mike", 'm', 18}
	// 方法值
	// f:=p.setInfoPointer
	fmt.Printf(("main:%p,%v\n"), &p, p)
	// 方法表达式
	f := (*Person).setInfoPointer
	f(&p)
}
