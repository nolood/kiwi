package cities

import (
	"fmt"
	"kiwi/.gen/kiwi/public/model"

	. "kiwi/.gen/kiwi/public/table"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repository interface {
	GetById(id int) (model.Cities, error)
}

type repository struct {
	log *zap.Logger
	db  *sqlx.DB
}

func New(log *zap.Logger, db *sqlx.DB) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}

func (r *repository) GetById(id int) (model.Cities, error) {
	const op = "repositories.cities.GetById"
	var city model.Cities

	stmt := SELECT(Cities.ID, Cities.Latitude, Cities.Longitude).FROM(Cities).WHERE(Cities.ID.EQ(Int(int64(id)))).LIMIT(1)

	err := stmt.Query(r.db, &city)
	if err != nil {
		return city, fmt.Errorf("%s: %w", op, err)
	}

	return city, nil
}
