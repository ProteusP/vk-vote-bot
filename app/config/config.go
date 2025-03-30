package config

import "os"

type Config struct {
	MattermostURL   string
	MattermostToken string
	TarantoolAddr   string
	WebhookToken    string
}

func Load() *Config {
	return &Config{
		MattermostURL:   os.Getenv("MATTERMOST_URL"),
		MattermostToken: os.Getenv("MATTERMOST_BOT_TOKEN"),
		TarantoolAddr:   os.Getenv("TARANTOOL_ADDR"),
		WebhookToken:    os.Getenv("WEBHOOK_TOKEN"),
	}
}
