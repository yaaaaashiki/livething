package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	returningHomeText    = "Welcome home!!!"
	illuminationInterval = 60
	wifiInterval         = 15
	alertRoopTimes       = 5
	zeroValue            = 0
)

type Slack struct {
	Text       string `json:"text"`
	Username   string `json:"username"`
	Icon_emoji string `json:"icon_emoji"`
	Icon_url   string `json:"icon_url"`
	Channel    string `json:"channel"`
}

var (
	IncomingUrl string = "https://hooks.slack.com/services/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"
)

func sendCurlRequest(text string) {
	params, _ := json.Marshal(Slack{
		fmt.Sprintf("%s", text),
		"livething",
		"",
		"",
		"#notification"})

	resp, _ := http.PostForm(
		IncomingUrl,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
}

func main() {
	sendCurlRequest("aaa")
}
