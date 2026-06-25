package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	go func() {
		<-timer.C //定时器阻塞
		fmt.Println("子携程可以打印了")
	}()
	// 停止
	timer.Stop()

	// 重置
	timer2 := time.NewTimer(3 * time.Second)
	ok := timer2.Reset(1 * time.Second)
	fmt.Println("ok=", ok)
	<-timer2.C
	fmt.Println("时间到")
}
