package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shenhailuanma/local-ip-notify/models"
	"github.com/shenhailuanma/local-ip-notify/service"
	"github.com/shenhailuanma/local-ip-notify/version"
	"os"
)

var (
	versionFlag = flag.Bool("v", false, "show version")
	configPath  = flag.String("c", "", "config file path")
)

func main() {
	flag.Parse()

	if *versionFlag {
		version.VersionInfo()
		os.Exit(0)
	}

	// load config
	if *configPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	configData, err := loadConfig(*configPath)
	if err != nil {
		fmt.Println("Failed to load config, error:", err.Error())
		os.Exit(1)
	}

	// fetch local ip
	localIP, err := service.FetchLocalIP()
	if err != nil {
		fmt.Println("Failed to fetch local IP, error:", err.Error())
		os.Exit(1)
	}

	// notify
	err = service.NotifyLocalIP(configData, localIP)
	if err != nil {
		fmt.Println("Failed to notify local IP, error:", err.Error())
		os.Exit(1)
	}
}

func loadConfig(configFilePath string) (models.Config, error) {
	var output = models.Config{}

	// read file data
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(data, &output)
	if err != nil {
		return output, err
	}

	// load env
	slackToken := os.Getenv("NotifySlackToken")
	slackChannel := os.Getenv("NotifySlackChannel")

	if slackToken != "" && slackChannel != "" {
		output.Slack = append(output.Slack, models.ConfigSlack{Token: slackToken, Channel: slackChannel})
	}

	return output, err
}
