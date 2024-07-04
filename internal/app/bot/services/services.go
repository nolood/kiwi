package services

import (
	"kiwi/internal/app/bot/repositories"
	"kiwi/internal/app/bot/services/profile"
	"kiwi/internal/app/bot/services/session"
	"kiwi/internal/app/bot/services/user"

	"go.uber.org/zap"
)

type Services struct {
	User    user.Service
	Session session.Service
	Profile profile.Service
}

func New(log *zap.Logger, repos *repositories.Repos) *Services {

	userService := user.New(log, repos)
	sessionService := session.New(log, repos)
	profileService := profile.New(log, repos)

	return &Services{
		User:    userService,
		Session: sessionService,
		Profile: profileService,
	}
}
