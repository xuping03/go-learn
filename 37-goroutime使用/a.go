package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is newTask")
		time.Sleep(time.Second) //延时一秒
	}
}
func main() {
	go newTask()
	for {
		fmt.Println("this is a main goroutime")
		time.Sleep(time.Second)
	}
}
