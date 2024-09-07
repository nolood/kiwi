package location

import (
	"encoding/json"
	"fmt"
	"kiwi/internal/app/meilisearch/constants"
	"kiwi/internal/app/meilisearch/dto"

	"github.com/meilisearch/meilisearch-go"
	"go.uber.org/zap"
)

type Service interface {
	Search(city string) (*CitySearchResponse, error)
}

type service struct {
	log    *zap.Logger
	client *meilisearch.Client
}

func New(log *zap.Logger, client *meilisearch.Client) Service {
	return &service{
		log:    log,
		client: client,
	}
}

type CitySearchResponse struct {
	meilisearch.SearchResponse
	Hits []dto.City `json:"hits"`
}

func (s *service) Search(city string) (*CitySearchResponse, error) {
	const op = "meilisearch.services.location.MustSearch"

	data, err := s.client.Index(constants.IndexCity).Search(city, &meilisearch.SearchRequest{
		Limit:               9,
		ShowMatchesPosition: true,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	hitsData, err := json.Marshal(data.Hits)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var result CitySearchResponse
	err = json.Unmarshal(hitsData, &result.Hits)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &result, nil
}
