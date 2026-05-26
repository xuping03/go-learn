package main

import (
	"encoding/json"
	"fmt"
)

// 成员变量首字母必须大写
type IT struct {
	Company  string
	Subjects []string
	IsOK     bool
	Price    float64
}

func main() {
	s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}
	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("buf=", string(buf))
	buf2, _ := json.MarshalIndent(s, "", " ")
	fmt.Println("buf2=", string(buf2))
}
