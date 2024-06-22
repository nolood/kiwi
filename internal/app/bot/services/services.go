package services

import (
	"kiwi/internal/app/bot/repositories"

	"go.uber.org/zap"
)

type Services struct {
	// ...
}

func New(log *zap.Logger, repos *repositories.Repos) *Services {
	return &Services{
		// ...
	}
}
