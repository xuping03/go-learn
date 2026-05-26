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
	jsonBuf := `
	 {
 "Company": "itcast",
 "Subjects": [
  "Go",
  "C++",
  "Python",
  "Test"
 ],
 "IsOK": true,
 "Price": 666.666
}
	`
	var tmp IT
	json.Unmarshal([]byte(jsonBuf), &tmp)
	fmt.Println("tmp", tmp)
	fmt.Printf("tmp=%+v\n", tmp)
}
