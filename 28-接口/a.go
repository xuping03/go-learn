package main

import "fmt"

type Humaner interface {
	sayHi()
}

type Student struct {
	name string
	age  int
}

func (tmp *Student) sayHi() {
	fmt.Println("student say hi", tmp)
}

type Teacher struct {
	name    string
	address string
}

func (tmp *Teacher) sayHi() {
	fmt.Println("teacher say hi", tmp)
}

type MyStr string

func (tmp *MyStr) sayHi() {
	fmt.Printf("mystr say hi %s\n", *tmp)
}

func main() {
	var i Humaner
	s := &Student{"mike", 6}
	i = s
	i.sayHi()

	t := &Teacher{"bj", "go"}
	i = t
	t.sayHi()

	var str MyStr = "hello mike"
	i = &str
	i.sayHi()
}
