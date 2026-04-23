package main

import (
	"fmt"
)

func main() {
	// 默认有break，需要fallthrough可以继续执行下面的case
	fmt.Println("hhh")
	goto end
end:
	fmt.Println("lll")
}
