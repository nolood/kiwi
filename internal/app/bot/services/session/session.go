package session

import (
	"errors"
	"fmt"
	"kiwi/.gen/kiwi/public/model"
	"kiwi/internal/app/bot/repositories"

	"github.com/go-jet/jet/v2/qrm"

	"go.uber.org/zap"
)

type Service interface {
	Get(tg_id int64) (model.Session, error)
	Set(tg_id int64, value model.Session) error
	Compare(tg_id int64, value model.Session) bool
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

func (s *service) Get(tg_id int64) (model.Session, error) {
	const op = "bot.services.session.Get"

	session, err := s.repos.Session.Get(tg_id)
	if err != nil {
		return session, fmt.Errorf("%s: %w", op, err)
	}

	return session, nil
}

func (s *service) Set(tg_id int64, value model.Session) error {
	const op = "bot.services.session.Set"

	err := s.repos.Session.Set(tg_id, value)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *service) Compare(tg_id int64, value model.Session) bool {
	const op = "bot.services.session.Compare"

	session, err := s.Get(tg_id)
	if err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			s.log.Error(op, zap.Error(err))
		}
		return false
	}

	if session == value {
		return true
	}

	return false
}
