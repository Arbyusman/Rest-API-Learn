package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {

	// Database
	DBHost     string `mapstructure:"DB_HOST"`
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_DBNAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	SSLMode    string `mapstructure:"SSL_MODE"`
	DBType     string `mapstructure:"DB_TYPE"`

	// app
	AppPort string `mapstructure:"APP_PORT"`

	// middlewares
	SecretJWT string `mapstructure:"SECRET_JWT"`

	// Cloudinary
	CloudName   string `mapstructure:"CLOUD_NAME"`
	CloudKey    string `mapstructure:"CLOUD_KEY"`
	ApiSecret   string `mapstructure:"API_SECRET"`
	CloudFolder string `mapstructure:"CLOUD_FOLDER"`
}

var (
	AppConfig Config
)

func LoadConfig() *Config {

	viper.SetConfigType("env")
	viper.SetConfigName("local")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return &AppConfig
}
