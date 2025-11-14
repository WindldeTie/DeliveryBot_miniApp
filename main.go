package main

import (
	"deliveryBot/handler"
	"deliveryBot/server"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	bot, err := initBot()
	if err != nil {
		log.Println(err)
	}
	go func() {
		handler.NewHandler(bot).Start(false)
	}()

	server.SetupServer()
}

func initBot() (*tgbotapi.BotAPI, error) {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("BOT_TOKEN environment variable is required")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		err = fmt.Errorf("bot creation failed: %v", err)
		return nil, err
	}

	log.Printf("✅ Authorized as @%s", bot.Self.UserName)
	return bot, nil
}
