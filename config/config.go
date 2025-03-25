package config

import "os"

type Config struct {
	MattermostURL   string
	MattermostToken string
	TarantoolAddr   string
}

func Load() *Config {
	return &Config{
		MattermostURL:   os.Getenv("MATTERMOST_URL"),
		MattermostToken: os.Getenv("MATTERMOST_TOKEN"),
		TarantoolAddr:   os.Getenv("TARANTOOL_ADDR"),
	}
}
