package user

import (
	"fmt"
	userdto "kiwi/internal/app/bot/dto/user"
	"kiwi/internal/app/bot/repositories"

	"github.com/mymmrac/telego"

	"go.uber.org/zap"
)

type Service interface {
	Get(tg_id int64) (userdto.UserWithProfile, error)
	Create(user *telego.User) (userdto.UserWithProfile, error)
	GetOrCreate(user *telego.User) (userdto.UserWithProfile, error)
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

func (s *service) Get(tg_id int64) (userdto.UserWithProfile, error) {
	const op = "bot.services.user.Get"

	user, err := s.repos.User.Get(tg_id)
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *service) Create(tgUser *telego.User) (userdto.UserWithProfile, error) {
	const op = "bot.services.user.Create"

	user, err := s.repos.User.Create(tgUser)
	if err != nil {
		s.log.Error(op, zap.Error(err))
		return user, err
	}

	return user, nil
}

func (s *service) GetOrCreate(tgUser *telego.User) (userdto.UserWithProfile, error) {
	const op = "bot.services.user.GetOrCreate"

	userprof, err := s.Get(tgUser.ID)
	if err != nil {
		return userprof, fmt.Errorf("%s: %w", op, err)
	}

	if userprof.User.ID != 0 {
		return userprof, nil
	}

	userprof, err = s.Create(tgUser)
	if err != nil {
		return userprof, fmt.Errorf("%s: %w", op, err)
	}

	return userprof, nil
}
