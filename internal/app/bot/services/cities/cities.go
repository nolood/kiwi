package cities

import (
	"fmt"
	"kiwi/.gen/kiwi/public/model"
	"kiwi/internal/app/bot/repositories"

	"go.uber.org/zap"
)

type Service interface {
	GetById(id int) (model.Cities, error)
	FindByCords(lat, lon float64) (model.Cities, float64, error)
}

type service struct {
	log   *zap.Logger
	repos *repositories.Repos
}

func New(log *zap.Logger, repos *repositories.Repos) Service {
	return &service{
		log:   log,
		repos: repos,
	}
}

func (s *service) GetById(id int) (model.Cities, error) {
	const op = "bot.services.cities.GetLocationById"

	city, err := s.repos.Cities.GetById(id)
	if err != nil {
		return city, fmt.Errorf("%s: %w", op, err)
	}

	return city, nil
}

func (s *service) FindByCords(lat, lon float64) (model.Cities, float64, error) {

	city, distance, err := s.repos.Cities.FindByCords(lat, lon)
	if err != nil {
		return city, distance, err
	}

	return city, distance, nil
}
