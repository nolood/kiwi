package services

import (
	"kiwi/internal/app/bot/repositories"
	"kiwi/internal/app/bot/services/user"

	"go.uber.org/zap"
)

type Services struct {
	User user.Service
}

func New(log *zap.Logger, repos *repositories.Repos) *Services {

	userService := user.New(log, repos)

	return &Services{
		User: userService,
	}
}
