package main

import (
	"log"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

func botCycle() {
	telegramToken := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}
	u := tgbotapi.NewUpdate(0)
	bot.GetUpdatesChan(u)

	bot.Debug = true
	channelName := os.Getenv("CHANNEL_NAME")
	log.Printf("Authorized on account %s", bot.Self.UserName)
	msg := tgbotapi.NewMessageToChannel(channelName, "text message")
	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
