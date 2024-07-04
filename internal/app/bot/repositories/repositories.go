package repositories

import (
	"kiwi/internal/app/bot/repositories/session"
	"kiwi/internal/app/bot/repositories/user"
	"kiwi/internal/storage/postgres"

	"go.uber.org/zap"
)

type Repos struct {
	User    user.Repository
	Session session.Repository
}

func New(log *zap.Logger, storage *postgres.Storage) *Repos {

	userRepo := user.New(log, storage.Db)
	sessionRepo := session.New(log, storage.Db)

	return &Repos{
		User:    userRepo,
		Session: sessionRepo,
	}
}
