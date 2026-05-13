package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

func (tmp Person) PrintInfo() {
	fmt.Println("tmp", tmp)
}

type long int

// 接受者类型不一样，这个方法计算同名，也是不同方法，不会出现重复定义函数的报错

func (tmp long) test() {

}

type char byte

func (tmp char) test() {

}

func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.age = a
	p.sex = s
}
func main() {
	p := Person{"张三", 20, 'm'}
	p.PrintInfo()
	p.SetInfo("李四", 22, 'f')
	p.PrintInfo()
}
