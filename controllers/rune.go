package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"unisat-telegram-bot/format"
	"unisat-telegram-bot/routes"
	"unisat-telegram-bot/types"
)

func GetRuneDetail(ticker string) (string, error) {
	url := fmt.Sprintf(routes.RunesDetailURL, ticker)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	authToken := os.Getenv("UNISAT_TOKEN")
	if authToken == "" {
		return "", fmt.Errorf("bearer token not set in environment")
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get data: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var runesDetailResponse types.RunesDetailResponse
	err = json.Unmarshal(body, &runesDetailResponse)
	if err != nil {
		return "", err
	}

	data := runesDetailResponse.Data
	return format.FormatRunesDetail(data), nil
}
