package profile

import (
	"kiwi/internal/app/bot/services"

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

func (s *Scene) RegisterFillProfileScene() {
	s.handleAge(s.GetGender)
	s.handleGender(s.GetPhoto)
	s.handleDefaultPhoto(s.GetAbout)
	s.handlePhoto(s.GetAbout)
	s.handleAbout(s.GetLocation)

	s.handleLocation(func(chatId telego.ChatID) {
		s.log.Info("kek")
	})

}
