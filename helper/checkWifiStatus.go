package helper

//If ping status is no problem, return true
//otherwise, return false
func checkWifiStatus() bool {
	res := PingStaticIP()
	if res == nil {
		return false
	}
	return true
}
