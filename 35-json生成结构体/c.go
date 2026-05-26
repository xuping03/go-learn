package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 4)
	m["c"] = "hh"
	m["s"] = []string{"Go", "C++", "python"}
	m["ok"] = true
	m["p"] = 55

	result, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("res", string(result))
}
