package config

import (
	"github.com/BurntSushi/toml"
	"fmt"
)

type Config struct {
	Notification NotificationConfig
	Mac          MacConfig
	Slack        SlackConfig
}

type NotificationConfig struct {
	RemindTitle           string `toml:"remind_title"`
	UpdatedTitle          string `toml:"updated_title"`
	PrefixNewSchedule     string `toml:"prefix_new_schedule"`
	PrefixDeletedSchedule string `toml:"prefix_deleted_schedule"`
	MinuteBefore          int    `toml:"minutes_before"`
	Interval              string `toml:"interval"`
}

type MacConfig struct {
	Enabled   bool `toml:"enabled"`
	SoundFlag bool `toml:"mute"`
}

type SlackConfig struct {
	Enabled  bool   `toml:"enabled"`
	Channel  string `toml:"channel"`
	Icon     string `toml:"icon_emoji"`
	UserName string `toml:"username"`
	Token    string `toml:"web_hook_token"`
}

var sharedInstance = newInstance()

func newInstance() *Config {
	var config Config
	if _, err := toml.DecodeFile("config.tml", &config); err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	return &config
}

func GetInstance() *Config {
	return sharedInstance
}
