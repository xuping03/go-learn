package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	i := 0
	for {
		<-ticker.C
		i++
		fmt.Println("i=", i)
		if i == 5 {
			ticker.Stop()
			break
		}
	}
}
