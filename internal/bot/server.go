package tgbot

import (
	"shekosaurus/internal/configs"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func Start(config *configs.Config) {
	logger := configureLogger(config)

	bot, err := createBot(config, logger)
	if err != nil {
		logger.Error(err)
	}

	updatesConfig := tgbotapi.NewUpdate(0)
	updatesConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updatesConfig)
	if err != nil {
		logger.Error(err)
	}

	Serve(bot, updates, logger)
}

func Serve(
	bot *tgbotapi.BotAPI,
	updates tgbotapi.UpdatesChannel,
	logger *logrus.Logger,
) {
	logger.Info("Serve")
	for update := range updates {
		logger.Info(update.Message)

		if update.Message != nil {
			if update.Message.Text != "" {
				handler(bot, update, logger)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я не понимаю :(")
				bot.Send(msg)
			}
		}
	}
}

func handler(bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
	logger *logrus.Logger,
) {
	switch update.Message.Text {
	case "/start":
		logger.Info("/start")

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, любимый щекозавр")
		bot.Send(msg)

	case "🦔":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что щекозавр?")
		bot.Send(msg)

	case "❤️":

	case "Ёжик":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что щекозавр?")
		bot.Send(msg)

	case "ёжик":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что щекозавр?")
		bot.Send(msg)

	case "/compliment":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Это пока не реализовано. Подожди")
		bot.Send(msg)

	default:
		logger.Info(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я такого не знаю(")
		bot.Send(msg)
	}
}
