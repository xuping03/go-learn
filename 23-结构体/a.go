package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	// 顺序初始化
	var s1 Student = Student{1, "xiaoming", 'm', 18, "bj"}
	fmt.Println("s1=", s1)
	// 指定成员初始化，没有初始化的成员，自动赋值为零值
	var s2 = Student{name: "xiaohong", addr: "sh"}
	fmt.Println("s2=", s2)

	// var p1 *Student = &Student{1, "xiaoming", 'm', 18, "bj"}
	// fmt.Println("*p1=", p1)

	p2 := &Student{name: "mike", addr: "sh"}
	fmt.Println("*p2=", p2)

	// 定义普通结构体
	var s Student
	// 定义指针变量，保存s的地址
	var p1 *Student
	p1 = &s
	// 通过指针操作成员 p1.id 和（*p1）.id完全等价
	p1.id = 1
	(*p1).name = "mike"
	p1.sex = 'f'
	p1.age = 20
	fmt.Println("s=", s)
	fmt.Println("p1=", p1)

	// 通过new 申请一个结构体
	p3 := new(Student)
	p3.id = 1
	p3.name = "linux"
	p3.sex = 'm'
	p3.age = 20
	p3.addr = "sz"
	fmt.Println("p3=", p3)

	// 如果想使用别的包的函数、结构体类型、结构体成员、函数名、类型名，结构体成员变量名，首字母必须大写可见
	// 如果首字母小写，只能在同一个包里使用
}
