package wifi

import (
	"time"
)

const (
	interval  = 5
	floatZero = 0.0
)

//Reference this varibale to check wifi status
var Status bool

//If ping status is no problem, return true
//Otherwise, return false
func CheckStatus() bool {
	res := PingStaticIP()
	if res.PacketLoss == floatZero {
		return true
	}
	return false
}

func SetCurrentStatus() {
	numberOfRoop := 0

	for {
		// For debug
		/*
			if numberOfRoop >= 5 {
				break
			}
		*/
		//fmt.Println(CheckStatus())
		Status = CheckStatus()
		if numberOfRoop != zeroValue {
			time.Sleep(interval * time.Second)
		}
		numberOfRoop++
	}
}
