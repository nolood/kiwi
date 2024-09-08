package profile

import (
	"fmt"
	userdto "kiwi/internal/app/bot/dto/user"
	"kiwi/internal/app/bot/repositories"
	"strconv"

	tu "github.com/mymmrac/telego/telegoutil"

	"github.com/mymmrac/telego"
	"go.uber.org/zap"
)

type Service interface {
	UpdateProfile(tg_id int64, profile userdto.ProfileUpdate) error
	GetFormattedProfile(chatId telego.ChatID) (*telego.SendPhotoParams, error)
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

func (s *service) GetFormattedProfile(chatId telego.ChatID) (*telego.SendPhotoParams, error) {
	const op = "services.profile.GetFormattedProfile"

	var text string

	userprof, err := s.repos.User.Get(chatId.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	text += *userprof.Profile.Name + "\n"

	text += strconv.Itoa(int(*userprof.Profile.Age)) + "\n"

	if userprof.Profile.About != nil {
		text += *userprof.Profile.About + "\n"
	}

	text += *userprof.Profile.Gender + "\n"

	msg := tu.Photo(chatId, telego.InputFile{FileID: *userprof.Profile.PhotoID})
	msg.Caption = text

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

	if profile.Name != nil {
		err := s.repos.User.UpdateName(tg_id, *profile.Name)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if profile.Location != nil {
		err := s.repos.User.UpdateCity(tg_id, *profile.Location)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}
