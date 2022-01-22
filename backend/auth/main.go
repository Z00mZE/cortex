package main

import (
	"github.com/Z00mZE/cortex/backend/auth/config"
	"github.com/Z00mZE/cortex/backend/auth/internal/app"
	"github.com/labstack/gommon/log"
)

func main() {
	//	получаем настройки
	configuration, configError := config.NewAuthConfig()
	//	что-ьл пошло не так, смысла продолжать нет
	if configError != nil {
		log.Fatalf("Config error: %s", configError)
	}
	//	запуск сервера приложения
	app.Run(configuration)
}
