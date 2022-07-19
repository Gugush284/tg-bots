package tgbot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

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
		logger.Info("Смайл ежа")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что щекозавр?")
		bot.Send(msg)

	case "❤️":
		logger.Info("Смайл сердца")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "❤️")
		bot.Send(msg)

	case "😘":
		logger.Info("Смайл поцелуй")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "И я тебя")
		bot.Send(msg)

	case "Ёжик":
		logger.Info("Ёжик")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что любимая?")
		bot.Send(msg)

	case "ёжик":
		logger.Info("ёжик")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я Ёжик")
		bot.Send(msg)

	case "/compliment":
		logger.Info("/compliment")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Это пока не реализовано. Подожди")
		bot.Send(msg)

	default:
		logger.Info(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я такого не знаю(")
		bot.Send(msg)
	}
}
