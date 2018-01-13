package usecase

import (
	"fmt"
	"time"

	"github.com/yaaaaashiki/livething/domain/slack"

	"log"
)

type NotificationUsecase struct {
	SlackAPIService *service.SlackAPIService
}

func NewNotificationUsecase(SlackAPIService *slack.SlackAPIService) *NotificationUsecase {
	return &NotificationUsecase{
		SlackAPIService: SlackAPIService,
	}
}

func (self *NotificationUsecase) Execute() {
	//hogehoge

	if err != nil {
		log.Println(err)
	}

		//fmt.Printf("[notification]%s [%s] is send!", time.Now().String(), hoge)
	}
}
