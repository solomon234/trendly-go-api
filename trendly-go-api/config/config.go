package config

import (
	"os"
)

type Config struct {
	// DB *DBConfig
	TikTokAPI    *APIConfig
	YoutubeAPI   *APIConfig
	TwitterAPI   *APIConfig
	InstagramAPI *APIConfig
}

type APIConfig struct {
	URL           string
	KEY           string
	HeaderAuthKey string
}

func GetConfig() *Config {
	return &Config{
		TikTokAPI: &APIConfig{
			URL:           os.Getenv("TIK_TOK_API_URL"),
			KEY:           os.Getenv("TIK_TOK_API_KEY"),
			HeaderAuthKey: "X-API-KEY",
		},
	}
}
