package service

import (
	"io/ioutil"
	"net/http"
)

func FetchLocalIP() (string, error) {
	const requestAPI = "https://ipinfo.io/ip"

	resp, err := http.Get(requestAPI)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
