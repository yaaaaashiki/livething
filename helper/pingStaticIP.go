package helper

import (
	"os"
	"os/exec"
)

const (
	ping     = "ping"
	staticIP = "127.0.0.1"
)

func PingStaticIP() {
	res := exec.Command(ping, staticIP)
	res.Stdout = os.Stdout
	res.Stderr = os.Stderr
	res.Run()
}
