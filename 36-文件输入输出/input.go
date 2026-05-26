package main

import "fmt"

func main() {
	// os.Stdout.Close(); //关闭后，无法输入

	fmt.Println("请输入a:")
	var a int
	fmt.Scan(&a)
}
