package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置种子，只需要一次
	// 如果种子参数一样，每次运行程序产生的随机数都一样
	rand.Seed(time.Now().UnixNano()) //以当前系统时间作为种子参数
	var a [10]int
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)
	// 冒泡排序
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)

}
