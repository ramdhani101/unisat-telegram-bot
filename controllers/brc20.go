package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"unisat-telegram-bot/format"
	"unisat-telegram-bot/routes"
)

func GetBRC20Detail(ticker string) (string, error) {
	url := fmt.Sprintf(routes.BRC20DetailURL, ticker)

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

	var brc20DetailResponse map[string]interface{}
	err = json.Unmarshal(body, &brc20DetailResponse)
	if err != nil {
		return "", err
	}

	data := brc20DetailResponse["data"].(map[string]interface{})
	return format.FormatBRC20Detail(data), nil
}
