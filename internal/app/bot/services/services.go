package services

import (
	"kiwi/internal/app/bot/repositories"
	"kiwi/internal/app/bot/services/cities"
	"kiwi/internal/app/bot/services/profile"
	"kiwi/internal/app/bot/services/session"
	"kiwi/internal/app/bot/services/user"
	"kiwi/internal/app/meilisearch"

	"go.uber.org/zap"
)

type Services struct {
	User    user.Service
	Session session.Service
	Profile profile.Service
	Cities  cities.Service

	MApp *meilisearch.App
}

func New(log *zap.Logger, repos *repositories.Repos, mApp *meilisearch.App) *Services {

	userService := user.New(log, repos)
	sessionService := session.New(log, repos)
	profileService := profile.New(log, repos)
	citiesService := cities.New(log, repos)

	return &Services{
		User:    userService,
		Session: sessionService,
		Profile: profileService,
		Cities:  citiesService,
		MApp:    mApp,
	}
}
