package main

import (
	"errors"
	"fmt"
	"kiwi/.gen/kiwi/public/model"
	. "kiwi/.gen/kiwi/public/table"
	m "kiwi/internal/app/meilisearch"
	"kiwi/internal/app/meilisearch/constants"
	"kiwi/internal/config"
	"kiwi/internal/lib/logger"
	"kiwi/internal/storage/postgres"
	"time"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/meilisearch/meilisearch-go"
	"go.uber.org/zap"
)

func main() {
	const op = "meilisearch.main"

	cfg := config.MustLoad()
	log := logger.New(cfg.Env)

	storage := postgres.New(cfg.Storage)

	meili := m.New(log, cfg.Meilisearch)

	const limit = 1000

	log.Info("Meilisearch: start migrating")

	// TODO: Вынести в сервис

	for offset := 0; offset <= 58000; offset += limit {
		stmt := SELECT(Cities.ID, Cities.Name, Cities.Alternatenames, Cities.Latitude, Cities.Longitude, Cities.CountryCode).FROM(Cities).OFFSET(int64(offset)).LIMIT(limit)

		var cities []model.Cities

		err := stmt.Query(storage.Db, &cities)
		if err != nil {
			if errors.Is(err, qrm.ErrNoRows) {
				log.Info(op, zap.String("offset", fmt.Sprintf("%d", offset)))
				break
			}

			log.Error(op, zap.Error(err))
		}

		if len(cities) == 0 {
			log.Info(op, zap.String("offset", fmt.Sprintf("%d", offset)))
			break
		}

		_, err = meili.Client.Index(constants.IndexCity).AddDocuments(cities)
		if err != nil {
			log.Error(op, zap.Error(err))
		} else {
			log.Info(fmt.Sprintf("Successfully indexed %d cities at offset = %d", len(cities), offset))
		}

		//checkStatus(meili, task.TaskUID, 0)

	}

	log.Info("Meilisearch: finish migrating")

}

func checkStatus(m *m.App, taskID int64, depth int) *meilisearch.Task {

	task, err := m.Client.GetTask(taskID)
	if err != nil {
		fmt.Println(err)
		return task
	}

	if task.Status == meilisearch.TaskStatusSucceeded {
		return task
	}

	fmt.Println(depth)

	time.Sleep(time.Second * 1)

	return checkStatus(m, taskID, depth+1)
}
