package scenes

import (
	"kiwi/internal/app/bot/handlers/scenes/profile"
	"kiwi/internal/app/bot/services"

	"github.com/mymmrac/telego"
	"go.uber.org/zap"
)

type Scenes struct {
	Profile *profile.Scene
}

func New(log *zap.Logger, servs *services.Services, bot *telego.Bot, updates <-chan telego.Update) *Scenes {

	profScene := profile.New(log, servs, bot, updates)

	return &Scenes{
		Profile: profScene,
	}
}
