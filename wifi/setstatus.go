package wifi

import (
	"fmt"
	"time"

	"github.com/yaaaaashiki/livething"
)

const (
	interval  = 5
	zeroValue = 0
)

//Reference this varibale to check wifi status

//If ping status is no problem, return true
//Otherwise, return false
func checkStatus() bool {
	res := sendPing()
	if res.PacketLoss == zeroValue {
		return true
	}
	return false
}

func SetCurrentStatus(wf *livething.Wifi) {
	numberOfRoop := 0

	for {
		fmt.Println(checkStatus())
		wf.Status = checkStatus()
		if numberOfRoop != zeroValue {
			time.Sleep(interval * time.Second)
		}
		numberOfRoop++
	}
}
