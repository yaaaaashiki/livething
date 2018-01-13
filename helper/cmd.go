package main

import (
	"fmt"
	"time"
)

const (
	interval  = 5
	zeroValue = 0
)

func main() {
	numberOfRoop := 0
	for {
		fmt.Println(checkWifiStatus())
		if numberOfRoop != zeroValue {
			time.Sleep(interval * time.Second)
		}
		numberOfRoop++
	}
}
