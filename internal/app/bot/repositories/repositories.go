package repositories

import (
	"kiwi/internal/app/bot/repositories/user"
	"kiwi/internal/storage/postgres"

	"go.uber.org/zap"
)

type Repos struct {
	User user.Repository
}

func New(log *zap.Logger, storage *postgres.Storage) *Repos {

	userRepo := user.New(log, storage.Db)

	return &Repos{
		User: userRepo,
	}
}
