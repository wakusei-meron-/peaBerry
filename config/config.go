package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Notification NotificationConfig
}

type NotificationConfig struct {
	UpdatedTitle          string `toml:"updated_title"`
	PrefixNewSchedule     string `toml:"prefix_new_schedule"`
	PrefixDeletedSchedule string `toml:"prefix_deleted_schedule"`
	SoundFlag             bool   `toml:"sound"`
	MinuteBefore          int    `toml:"minutes_before"`
}

var sharedInstance = newInstance()

func newInstance() *Config {
	var config Config
	if _, err := toml.DecodeFile("config.tml", &config); err != nil {
		panic(err)
	}

	return &config
}

func GetInstance() *Config {
	return sharedInstance
}
