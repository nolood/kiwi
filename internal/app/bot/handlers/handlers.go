package handlers

import (
	"kiwi/internal/app/bot/handlers/callbacks"
	"kiwi/internal/app/bot/handlers/commands"
	"kiwi/internal/app/bot/handlers/scenes"
	"kiwi/internal/app/bot/services"
	"sync"

	"github.com/mymmrac/telego"

	"go.uber.org/zap"
)

func Register(log *zap.Logger, updates <-chan telego.Update, b *telego.Bot, servs *services.Services) {
	sc := scenes.New(log, servs, b, updates)

	comms := commands.New(log, servs, b, updates)
	comms.Start()

	cb := callbacks.New(log, servs, b, updates, sc)
	cb.FillProfile()

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		comms.Bh.Start()
	}()

	// go func() {
	// defer wg.Done()
	// cb.Bh.Start()
	// }()

	for update := range updates {
		// Здесь обрабатываем каждое обновление, которое приходит из канала updates
		log.Info("Received update", zap.Any("update", update))

	}

	wg.Wait()

}
