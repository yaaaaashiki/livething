package main

import (
	"fmt"

	ping "github.com/sparrc/go-ping"
)

const (
	staticIP = "127.0.0.2"
)

//PingStaticIP
func main() {
	/*
		res := exec.Command(ping, staticIP)
		res.Stdout = os.Stdout
		res.Stderr = os.Stderr
		res.Run()
	*/

	pinger, err := ping.NewPinger(staticIP)
	if err != nil {
		fmt.Fprintf(err)
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
}
