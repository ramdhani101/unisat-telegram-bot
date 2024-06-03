package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unisat-telegram-bot/controllers"
	"unisat-telegram-bot/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN tidak ditemukan")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.Request(tgbotapi.DeleteWebhookConfig{})
	if err != nil {
		log.Fatalf("Gagal menghapus webhook: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		var msgText string
		switch {
		case strings.HasPrefix(update.Message.Text, "/start"):
			msgText = messages.StartMessage
		case strings.HasPrefix(update.Message.Text, "/brc20"):
			args := strings.Split(update.Message.Text, " ")
			if len(args) == 1 {
				msgText = messages.BRC20Message
			} else if len(args) == 2 {
				ticker := args[1]
				data, err := controllers.GetBRC20Detail(ticker)
				if err != nil {
					msgText = fmt.Sprintf("Failed to get details for %s", ticker)
				} else {
					msgText = data
				}
			} else {
				msgText = "Usage: /brc20 or /brc20 [ticker]"
			}
		case strings.HasPrefix(update.Message.Text, "/runes"):
			args := strings.Split(update.Message.Text, " ")
			if len(args) == 1 {
				msgText = messages.RuneMessage
			} else if len(args) == 2 {
				runeId := args[1]
				data, err := controllers.GetRuneDetail(runeId)
				if err != nil {
					msgText = fmt.Sprintf("Failed to get details for %s", runeId)
				} else {
					msgText = data
				}
			} else {
				msgText = "Usage: /runes or /runes [runeId]"
			}
		default:
			msgText = "Unknown command"
		}

		if msgText != "" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			bot.Send(msg)
		} else {
			log.Printf("Empty message text for command: %s", update.Message.Text)
		}

	}
}
