package user

import (
	"fmt"
	"kiwi/.gen/kiwi/public/model"
	"kiwi/internal/app/bot/repositories"

	"github.com/mymmrac/telego"

	"go.uber.org/zap"
)

type Service interface {
	Get(tg_id int64) (model.Users, error)
	Create(user *telego.User) (model.Users, error)

	GetOrCreate(user *telego.User) (model.Users, error)
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

func (s *service) Get(tg_id int64) (model.Users, error) {
	const op = "services.user.Get"

	user, err := s.repos.User.Get(tg_id)
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *service) Create(tgUser *telego.User) (model.Users, error) {
	const op = "services.user.Create"

	user, err := s.repos.User.Create(tgUser)
	if err != nil {
		s.log.Error(op, zap.Error(err))
		return user, err
	}

	return user, nil
}

func (s *service) GetOrCreate(tgUser *telego.User) (model.Users, error) {
	const op = "services.user.GetOrCreate"

	user, err := s.Get(tgUser.ID)
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	if user.ID != 0 {
		return user, nil
	}

	user, err = s.Create(tgUser)
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
