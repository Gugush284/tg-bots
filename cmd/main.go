package main

import (
	"flag"
	"log"

	"shekosaurus/internal/configs"

	"github.com/BurntSushi/toml"
)

var configPath string
var sessionKeyPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	_, err = toml.DecodeFile(sessionKeyPath, config)
	if err != nil {
		log.Fatal(err)
	}
}
