package config

import (
	"context"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Env         string `mapstructure:"env"`
	StoragePath string `mapstructure:"storage_path"`
	HTTPServer  `mapstructure:"http_server"`
}

type HTTPServer struct {
	Address     string        `mapstructure:"address"`
	Timeout     time.Duration `mapstructure:"timeout"`
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`
}

func LoadConfig(ctx context.Context) *Config {
	var config Config
	configPath, exists := os.LookupEnv("CONFIG_PATH")
	if exists {
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath("configs")
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("couldn't read config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("couldn't unmarshal config file: %v", err)
	}
	return &config
}
