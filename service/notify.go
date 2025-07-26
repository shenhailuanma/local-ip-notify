package service

import (
	"fmt"
	"github.com/shenhailuanma/local-ip-notify/models"
)

func NotifyLocalIP(configData models.Config, localIP string) error {

	msg := fmt.Sprintf("Local IP: %s", localIP)

	// slack notify
	for _, slackNotifyOne := range configData.Slack {
		err := NotifySlack(slackNotifyOne, msg)
		if err != nil {
			fmt.Printf("Failed to notify slack channel:%s, failed:%s", slackNotifyOne.Channel, err.Error())
			continue
		}
	}

	return nil
}
