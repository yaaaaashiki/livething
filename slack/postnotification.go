package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/yaaaaashiki/livething/model"
)

const (
	channelName          = "#notification"
	userName             = "livething"
	returningHomeText    = "Welcome home!!!"
	iconName             = "warning"
	illuminationInterval = 4
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

func setAlertText(objectName string) string {
	alertText := "Put" + objectName + "on the home position"
	return alertText
}

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

func PostNotification(object *model.Object, wifi *model.Wifi) {
	roopTimesCounter := zeroValue
	consecutiveCounter := zeroValue
	alertText := ""
	for {
		fmt.Printf("object status: ")
		fmt.Println(object.Status)
		fmt.Println(object.Name)
		if wifi.Status == false {
			time.Sleep(wifiInterval * time.Second)
			continue
		}

		if roopTimesCounter >= alertRoopTimes && consecutiveCounter != zeroValue && object.Status == false {
			//fmt.Println(object.Status)
			alertText = setAlertText(object.Name)
			sendCurlRequest(alertText)
			time.Sleep(illuminationInterval * time.Second)
			roopTimesCounter = zeroValue
			continue
		}

		if roopTimesCounter == zeroValue && consecutiveCounter == zeroValue {
			sendCurlRequest(returningHomeText)
			time.Sleep(illuminationInterval * time.Second)
			roopTimesCounter++
			continue
		}

		if object.Status == false {
			time.Sleep(illuminationInterval * time.Second)
			roopTimesCounter++
			consecutiveCounter++
		}
	}
}
