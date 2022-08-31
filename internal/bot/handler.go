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
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!")
		bot.Send(msg)

	case "/compliment":
		logger.Info("/compliment")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, complimentApi(logger))
		bot.Send(msg)

	default:
		logger.Info(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я такого не знаю(")
		bot.Send(msg)
	}
	logger.Info("Finished processing the request")
}
