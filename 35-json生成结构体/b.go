package main

import (
	"encoding/json"
	"fmt"
)

// 成员变量首字母必须大写
type IT struct {
	Company  string   `json:"-"`        //此字段不会输出到屏幕
	Subjects []string `json:"subjects"` //二次编码
	IsOK     bool     `json:",string"`  //以字符串的形式
	Price    float64  `json:",string"`
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
