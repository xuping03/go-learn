package main

import (
	"fmt"
)

func main() {
	// 默认有break，需要fallthrough可以继续执行下面的case
	var num int
	fmt.Printf("按下楼层：")
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println("你按下了一楼")
		fallthrough
	case 2:
		fmt.Println("你按下了二楼")
		fallthrough
	case 3:
		fmt.Println("你按下了三楼")
	case 4:
		fmt.Println("你按下了四楼")
	default:
		fmt.Println("你按下了其他楼层")
	}
	switch c := 1; c {
	case 1:
		fmt.Println("你按下了一楼")
	case 2:
		fmt.Println("你按下了二楼")
	case 3:
		fmt.Println("你按下了三楼")
	case 4:
		fmt.Println("你按下了四楼")
	default:
		fmt.Println("你按下了其他楼层")
	}
	var score = 78
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 70:
		fmt.Println("中等")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}
