package main

import (
	"flag"
	"log"

	tgserver "shekosaurus/internal/bot"
	"shekosaurus/internal/configs"

	"github.com/BurntSushi/toml"
)

var configPath string
var sessionKeyPath string

func init() {
	flag.StringVar(&configPath, "config-path", "internal/configs/tgbot.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	tgserver.Start(config)
}
