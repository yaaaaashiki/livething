package slack

import (
	"time"

	"github.com/yaaaaashiki/livething/interfaceadapter/controller"
	"github.com/yaaaaashiki/livething/wifi"
)

const (
	channelName               = "#notification"
	userName                  = "livething"
	returningHomeText         = "Welcome home!!!"
	iconName                  = "warning"
	postURI                   = "https://hooks.slack.com/services/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"
	checkIlluminationInterval = 60
	checkWifiInterval
	zeroValue      = 0
	alertRoopTimes = 5
	firstTime      = 1
)

func setAlertText(objectName string) (alertText string) {
	alertText = "Put" + objectName + "on the home position"
	return
}

func PostNotification() {

	roopTimesCounter := zeroValue

	for {
		if wifi.Status == false {
			time.Sleep(checkWifiInterval * time.Second)
			continue
		}

		if roopTimesCounter >= alertRoopTimes {
			//post transaction to slack
			roopTimesCounter = zeroValue
			continue
		}

		if roopTimesCounter == firstTime {
			//post welcome notiridatin to slack
			continue
		}

		if controller.ObjectStatus == false {
			roopTimesCounter++
		}
	}

	/*
		curl -X POST --data-urlencode "payload={\"channel\": \"#notification\", \"username\": \"livething\", \"text\": \"This is posted to #notification and comes from a bot named webhookbot.\", \"icon_emoji\": \":warning:\"}"
	*/
}
