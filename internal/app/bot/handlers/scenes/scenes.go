package scenes

import (
	"kiwi/internal/app/bot/handlers/scenes/profile"
	"kiwi/internal/app/bot/services"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"go.uber.org/zap"
)

type Scenes struct {
	Profile *profile.Scene
}

func New(log *zap.Logger, servs *services.Services, bot *telego.Bot, bh *th.BotHandler) *Scenes {

	profScene := profile.New(log, servs, bot, bh)

	return &Scenes{
		Profile: profScene,
	}
}

func (s *Scenes) Register() {
	s.Profile.RegisterFillProfileScene()
}
