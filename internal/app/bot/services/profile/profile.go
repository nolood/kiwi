package profile

import (
	"fmt"
	userdto "kiwi/internal/app/bot/dto/user"
	"kiwi/internal/app/bot/repositories"

	"github.com/mymmrac/telego"

	"go.uber.org/zap"
)

type Service interface {
	UpdateProfile(tg_id int64, profile userdto.ProfileUpdate) error
	GetFormattedProfile(chatId telego.ChatID) (*telego.SendMessageParams, error)
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

func (s *service) GetFormattedProfile(chatId telego.ChatID) (*telego.SendMessageParams, error) {
	const op = "services.profile.GetFormattedProfile"

	var msg *telego.SendMessageParams

	// userprof, err := s.repos.User.Get(chatId.ID)
	// if err != nil {
	// 	return msg, fmt.Errorf("%s: %w", op, err)
	// }

	// msg = tu.Message(chatId)

	return msg, nil

}

func (s *service) UpdateProfile(tg_id int64, profile userdto.ProfileUpdate) error {
	const op = "bot.services.profile.UpdateProfile"

	if profile.Age != nil {
		err := s.repos.User.UpdateAge(tg_id, *profile.Age)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if profile.Gender != nil {
		err := s.repos.User.UpdateGender(tg_id, *profile.Gender)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if profile.About != nil {
		err := s.repos.User.UpdateAbout(tg_id, *profile.About)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if profile.PhotoId != nil {
		err := s.repos.User.UpdatePhoto(tg_id, *profile.PhotoId)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if profile.Latitude != nil && profile.Longitude != nil {
		err := s.repos.User.UpdateLocation(tg_id, profile.Latitude, profile.Longitude)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}
