package wifi

import (
	"time"
)

const (
	interval  = 5
	zeroValue = 0
)

var Status bool

func getCurrentStatus() {
	numberOfRoop := 0

	for {
		if numberOfRoop >= 5 {
			break
		}
		//fmt.Println(checkWifiStatus())
		Status = checkWifiStatus()
		if numberOfRoop != zeroValue {
			time.Sleep(interval * time.Second)
		}
		numberOfRoop++
	}
}
