package wifi

import (
	"fmt"
	"log"

	ping "github.com/sparrc/go-ping"
)

const (
	staticIP = "127.0.0.1"
)

func PingStaticIP() *ping.Statistics {
	pinger, err := ping.NewPinger(staticIP)
	if err != nil {
		log.Println(err)
	}
	pinger.Count = 3
	pinger.Run()               // blocks until finished
	res := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(res)
	return res
}
