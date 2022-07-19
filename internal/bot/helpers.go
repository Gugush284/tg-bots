package tgbot

import (
	"log"
	"shekosaurus/internal/configs"

	tgbotapi "github.com/Syfaro/telegram-bot-api"

	"github.com/sirupsen/logrus"
)

func configureLogger(config *configs.Config) *logrus.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	logger.SetLevel(level)

	return logger
}

func createBot(config *configs.Config, logger *logrus.Logger) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.Apitoken)
	if err != nil {
		return nil, err
	}

	bot.Debug = false

	logger.Infof("Authorized on account %s", bot.Self.UserName)

	return bot, nil
}
