package profile

import (
	"kiwi/.gen/kiwi/public/model"
	"kiwi/internal/app/bot/services"
	"strings"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"go.uber.org/zap"
)

type Scene struct {
	log      *zap.Logger
	services *services.Services
	bh       *th.BotHandler
	bot      *telego.Bot
}

func New(log *zap.Logger, servs *services.Services, bot *telego.Bot, bh *th.BotHandler) *Scene {

	return &Scene{
		log:      log,
		services: servs,
		bh:       bh,
		bot:      bot,
	}
}

func (s *Scene) router(next func(chatId telego.ChatID, session interface{})) func(chatId telego.ChatID, session model.Session) {

	return func(chatId telego.ChatID, session model.Session) {
		if strings.Split(session.String(), "_")[0] == "edit" {
			s.GetProfileComplete(chatId, nil)
			return
		}
		next(chatId, nil)
	}

}

func (s *Scene) RegisterFillProfileScene() {
	s.handleAge(s.router(s.GetGender))
	s.handleGender(s.router(s.GetPhoto))
	s.handleDefaultPhoto(s.router(s.GetAbout))
	s.handlePhoto(s.router(s.GetAbout))
	s.handleAbout(s.router(s.GetLocation))
	s.handleLocation(s.router(s.GetProfileComplete))
	s.handleLocationTown(s.router(s.GetProfileComplete))

	s.handleEditProfile(s.router(s.GetProfileComplete))
}
