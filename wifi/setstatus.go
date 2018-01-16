package wifi

import (
	"time"

	"github.com/yaaaaashiki/livething/model"
)

const (
	interval  = 5
	zeroValue = 0
)

//If ping status is no problem, return true
//Otherwise, return false
func checkStatus() bool {
	res := sendPing()
	if res.PacketLoss == zeroValue {
		return true
	}
	return false
}

func SetCurrentStatus(wf *model.Wifi) {
	numberOfRoop := 0

	for {
		wf.Status = checkStatus()
		if numberOfRoop != zeroValue {
			time.Sleep(interval * time.Second)
		}
		numberOfRoop++
	}
}
