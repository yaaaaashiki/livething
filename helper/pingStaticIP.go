package helper

import (
	"fmt"
	"os/exec"
)

const (
	staticIP = "127.0.0.1"
)

func PingStaticIP() {
	out, err := exec.Command("ping", staticIP).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(out))
	fmt.Println("hoge")
}
