package repositories

import (
	"kiwi/internal/app/bot/repositories/cities"
	"kiwi/internal/app/bot/repositories/session"
	"kiwi/internal/app/bot/repositories/user"
	"kiwi/internal/storage/postgres"

	"go.uber.org/zap"
)

type Repos struct {
	User    user.Repository
	Session session.Repository
	Cities  cities.Repository
}

func New(log *zap.Logger, storage *postgres.Storage) *Repos {

	userRepo := user.New(log, storage.Db)
	sessionRepo := session.New(log, storage.Db)
	citiesRepo := cities.New(log, storage.Db)

	return &Repos{
		User:    userRepo,
		Session: sessionRepo,
		Cities:  citiesRepo,
	}
}
