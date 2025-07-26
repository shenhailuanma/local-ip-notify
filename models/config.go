package models

type Config struct {
	Slack []ConfigSlack `json:"slack"`
}

type ConfigSlack struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
}
