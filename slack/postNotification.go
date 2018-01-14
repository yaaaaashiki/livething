package slack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/yaaaaashiki/livething/interfaceadapter/controller"
	"github.com/yaaaaashiki/livething/wifi"
)

const (
	channelName          = "#notification"
	userName             = "livething"
	returningHomeText    = "Welcome home!!!"
	iconName             = "warning"
	illuminationInterval = 60
	wifiInterval         = 15
	zeroValue            = 0
	firstTime            = 0
	alertRoopTimes       = 5
)

var alertText string

type Slack struct {
	text        string `json:"text"`
	userName    string `json:"username"`
	iconName    string `json:"icon_emoji"`
	channelName string `json:"channel"`
}

func setAlertText(objectName string) string {
	alertText = "Put" + objectName + "on the home position"
	return alertText
}

func curlRequest(text string) {
	params, _ := json.Marshal(Slack{
		text:        text,
		username:    userName,
		iconName:    iconName,
		iconURI:     "",
		channelName: channelName})

	resp, _ := http.PostForm(
		os.Getenv("WEBHOOKURI"),
		url.Values{"payload": {string(params)}},
	)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
}

func PostNotification() {
	roopTimesCounter := zeroValue

	for {
		object := &controller.Object{}

		if wifi.Status == false {
			time.Sleep(wifiInterval * time.Second)
			continue
		}

		if roopTimesCounter >= alertRoopTimes {
			setAlertText(object.Name)
			curlRequest(alertText)
			time.Sleep(illuminationInterval * time.Second)
			roopTimesCounter = zeroValue
			continue
		}

		if roopTimesCounter == firstTime {
			curlRequest(returningHomeText)
			time.Sleep(illuminationInterval * time.Second)
			roopTimesCounter++
			continue
		}

		if object.Status == false {
			time.Sleep(illuminationInterval * time.Second)
			roopTimesCounter++
		}
	}
}
