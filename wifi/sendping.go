package wifi

import (
	"log"

	ping "github.com/sparrc/go-ping"
)

const (
	staticIP      = "133.2.210.50"
	sendPingTimes = 3
)

func sendPing() *ping.Statistics {
	pinger, err := ping.NewPinger(staticIP)
	if err != nil {
		log.Println(err)
	}

	pinger.Count = sendPingTimes
	pinger.Run()               // blocks until finished
	res := pinger.Statistics() // get send/receive/rtt stats
	return res
}
