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
		if value, ok := data.(int); ok == true {
			fmt.Println(value, ok, index)
		}
	}
}
