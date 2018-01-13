package wifi

//If ping status is no problem, return true
//Otherwise, return false
func CheckStatus() bool {
	res := PingStaticIP()
	if res == nil {
		return false
	}
	return true
}
