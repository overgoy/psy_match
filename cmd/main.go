package main

import (
	"context"
	"log"
	"psy_match/config"
	"psy_match/internal/bot"
	"psy_match/internal/database"
)

func main() {

	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Ошибка при загрузке конфигурации: %v", err)
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close(context.Background())

	log.Println("Starting Telegram bot...")
	bot.StartBot()
}
