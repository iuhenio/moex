package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func sendMessage(text string, c config) {
	bot, err := tgbotapi.NewBotAPI(c.TlgAccessToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	//fmt.Println("Authorized on account", bot.Self.UserName)

	msg := tgbotapi.NewMessage(c.TlgChatId, text)

	bot.Send(msg)
}
