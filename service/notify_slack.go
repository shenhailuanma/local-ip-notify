package service

import (
	"errors"
	"github.com/shenhailuanma/local-ip-notify/models"
	"github.com/slack-go/slack"
)

func NotifySlack(configSlack models.ConfigSlack, message string) error {
	// check
	if configSlack.Token == "" {
		return errors.New("empty token")
	}
	if configSlack.Channel == "" {
		return errors.New("empty channel")
	}
	if message == "" {
		return errors.New("empty message")
	}

	api := slack.New(configSlack.Token)
	_, _, err := api.PostMessage(configSlack.Channel, slack.MsgOptionText(message, false))
	return err
}
