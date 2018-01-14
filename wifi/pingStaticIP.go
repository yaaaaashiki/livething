package wifi

import (
	"fmt"
	"log"

	ping "github.com/sparrc/go-ping"
)

const (
	staticIP = "133.2.210.50"
)

func PingStaticIP() *ping.Statistics {
	pinger, err := ping.NewPinger(staticIP)
	if err != nil {
		log.Println(err)
	}

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	pinger.Count = 3
	pinger.Run()               // blocks until finished
	res := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(res)
	return res
}
