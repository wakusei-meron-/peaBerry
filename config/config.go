package config

type Config struct {
	Notification NotificationConfig
}

type NotificationConfig struct {
	UpdatedTitle          string `toml:"updated_title"`
	PrefixNewSchedule     string `toml:"prefix_new_schedule"`
	PrefixDeletedSchedule string `toml:"prefix_deleted_schedule"`
	SoundFlag             bool   `toml:"sound"`
}
