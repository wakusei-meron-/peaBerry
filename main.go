package main

import (
	"peaberry/domain"
	"github.com/BurntSushi/toml"
	"peaberry/config"
)

func main() {
	var config config.Config
	if _, err := toml.DecodeFile("config.tml", &config); err != nil {
		panic(err)
	}

	domain.StartApplication(config)
}

