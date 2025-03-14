package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	BotToken    string `mapstructure:"bot_token"` // Используется bot_token в конфигурации
	DatabaseURL string `mapstructure:"database_url"`
}

func LoadConfig() error {
	// Указываем путь к конфигурационным файлам и тип конфигурации
	viper.AddConfigPath("./config") // Текущая директория
	viper.SetConfigName("config")   // Имя конфигурационного файла
	viper.SetConfigType("json")     // Формат конфигурации

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("ошибка при чтении конфигурации: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("ошибка при маппинге конфигурации: %w", err)
	}

	AppConfig = &cfg
	return nil
}
