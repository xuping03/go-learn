package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getNum(arr []int, num int) {
	arr[0] = num / 1000
	arr[1] = num / 100 % 10
	arr[2] = num / 10 % 10
	arr[3] = num % 10
}

func getInputNum() int {
	var inputNum int
	fmt.Println("请输入一个数字：")
	for {
		fmt.Scan(&inputNum)
		if inputNum > 999 && inputNum < 10000 {
			break
		}
	}

	return inputNum
}

func guess(inputSlice []int, randSlice []int) {
	n := 0
	for i := 0; i < 4; i++ {
		if randSlice[i] == inputSlice[i] {
			fmt.Printf("第%d位正确\n", i+1)
			n++
		} else if randSlice[i] > inputSlice[i] {
			fmt.Printf("第%d位太小\n", i+1)
		} else if randSlice[i] < inputSlice[i] {
			fmt.Printf("第%d位太大\n", i+1)
		}
	}
	if n == 4 {
		fmt.Println("恭喜你，猜对了！")
	}
}

// 数组作为函数参数，是值传递

func main() {
	rand.Seed(time.Now().UnixNano())
	var randNum int
	for {
		randNum = rand.Intn(10000)
		fmt.Println(randNum)
		if randNum > 999 {
			break
		}
	}
	randSlice := make([]int, 4)
	getNum(randSlice, randNum)
	inputNum := getInputNum()
	inputSlice := make([]int, 4)
	getNum(inputSlice, inputNum)

	guess(inputSlice, randSlice)
}
