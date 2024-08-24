package services

import (
	"kiwi/internal/app/meilisearch/services/location"

	"github.com/meilisearch/meilisearch-go"
	"go.uber.org/zap"
)

type Services struct {
	Location location.Service
}

func New(log *zap.Logger, client *meilisearch.Client) *Services {

	locationService := location.New(log, client)

	return &Services{
		Location: locationService,
	}
}
