package configs

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type EnvironmentConfig struct {
	Opensearch OpenSearchConfig
}

type OpenSearchConfig struct {
	Host           string `envconfig:"OPEN_SEARCH_HOST"`
	VideoDataIndex string `envconfig:"OPEN_SEARCH_VIDEO_DATA_INDEX"`
	Username       string `envconfig:"OPEN_SEARCH_USERNAME"`
	Password       string `envconfig:"OPEN_SEARCH_PASSWORD"`
}

var config EnvironmentConfig

func GetEnvironmentConfig() *EnvironmentConfig {
	if config != (EnvironmentConfig{}) {
		return &config
	}

	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal("Could not load Configs: ", err)
	}

	return &config
}
