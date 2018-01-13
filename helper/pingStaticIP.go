package helper

import (
	"fmt"
	"log"

	ping "github.com/sparrc/go-ping"
)

const (
	staticIP = "127.0.0.1"
)

func PingStaticIP() *ping.Statistics {
	/*
		res := exec.Command(ping, staticIP)
		res.Stdout = os.Stdout
		res.Stderr = os.Stderr
		res.Run()
	*/

	pinger, err := ping.NewPinger(staticIP)
	if err != nil {
		log.Println(err)
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	res := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(res)
	return res
}
