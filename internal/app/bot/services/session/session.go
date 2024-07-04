package session

import (
	"kiwi/internal/app/bot/repositories"

	"go.uber.org/zap"
)

type Service interface {
	Get(tg_id int64) (string, error)
	Set(tg_id int64, value string) (string, error)
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

func (s *service) Get(tg_id int64) (string, error) {
	const op = "services.session.Get"

	s.log.Info(op)

	return "", nil
}

func (s *service) Set(tg_id int64, value string) (string, error) {
	const op = "services.session.Set"

	s.log.Info(op)

	return "", nil
}
