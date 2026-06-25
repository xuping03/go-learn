package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)
	// 换冲突剩余个数，缓冲大小
	fmt.Println("len(ch)", len(ch), cap(ch))
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子携程", i)
			ch <- i
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println("num=", num)
	}
}
