package app

import (
	botapp "kiwi/internal/app/bot"
	"kiwi/internal/app/bot/repositories"
	"kiwi/internal/app/bot/services"
	"kiwi/internal/config"
	"kiwi/internal/storage/postgres"

	"go.uber.org/zap"
)

type App struct {
	Bot *botapp.BotApp
}

func New(log *zap.Logger, cfg *config.Config) *App {

	storage := postgres.New(cfg.Storage)

	repos := repositories.New(log, storage)

	services := services.New(log, repos)

	bot := botapp.New(log, cfg.Telegram, services)

	return &App{
		Bot: bot,
	}
}
