package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	telegramToken := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	channelName := os.Getenv("CHANNEL_NAME")
	log.Printf("Authorized on account %s", bot.Self.UserName)
	msg := tgbotapi.NewMessageToChannel(channelName, "text message")
	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
