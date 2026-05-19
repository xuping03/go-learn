package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Student{"mike", 666}
	// 类型查询，类型断言

	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Println("int", value, index)
		case string:
			fmt.Println("string", value, index)

		case Student:
			fmt.Println("Student", value, index)
		}
	}
}
