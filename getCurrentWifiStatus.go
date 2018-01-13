package helper

import (
	"time"
)

const (
	interval   = 5
	zeroValue  = 0
	WifiStatus = true
)

func getCurrentWifiStatus() {
	numberOfRoop := 0
	for {
		//fmt.Println(checkWifiStatus())

		if numberOfRoop != zeroValue {
			time.Sleep(interval * time.Second)
		}
		numberOfRoop++
	}
}
