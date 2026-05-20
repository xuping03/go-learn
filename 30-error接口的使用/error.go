package main

import "fmt"

func main() {
	err1 := fmt.Errorf("%s", "this is normal err1")
	fmt.Println("err1=", err1)

}
