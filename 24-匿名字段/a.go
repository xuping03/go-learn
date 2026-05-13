package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	school string
	id     int
}

type mystr string
type Student2 struct {
	Person
	int
	mystr
}

func main() {
	var s1 Student = Student{Person{"mike", 20}, "qinghua", 123}
	fmt.Println("s1=", s1)

	s2 := Student{Person{"mike", 20}, "qinghua", 123}
	fmt.Println("s2=", s2)
	// %+v，显示更详细
	fmt.Printf("s2=%+v\n", s2)

	s3 := Student{id: 1}
	fmt.Printf("s3=%+v\n", s3)

	s4 := Student{Person: Person{name: "mike"}}
	fmt.Printf("s4=%+v\n", s4)

	s5 := Student2{Person{name: "xx"}, 12, "tt"}
	fmt.Printf("s5=%+v\n", s5)
}
