package main //helper

//If ping status is no problem, return true
//Otherwise, return false

func checkWifiStatus() bool {
	res := PingStaticIP()
	if res == nil {
		return false
	}
	return true
}
