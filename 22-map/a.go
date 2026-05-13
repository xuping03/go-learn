package main

import "fmt"

func main() {
	// 键值类型：int ，value类型：字符串
	var m1 map[int]string
	fmt.Println("m1", m1)

	m2 := make(map[int]string)
	fmt.Println("m2", m2)

	m3 := make(map[int]string, 2)
	m3[1] = "a"
	m3[2] = "b"
	m3[3] = "c"

	fmt.Println("m3", m3)

	// map的遍历
	for key, value := range m3 {
		fmt.Printf("%d====>%s\n", key, value)
	}

	value, ok := m3[1]
	if ok == true {
		fmt.Println("m3[1]=", value)
	} else {
		fmt.Println("m3[1]不存在")
	}
	test(m3)
	fmt.Println("m3", m3)
}
func test(m map[int]string) {
	delete(m, 1)
}
