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
	GetByCords(lat, lon float64) (model.Cities, error)
	FindByCords(lat, lon float64) (model.Cities, float64, error)
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

func (r *repository) GetByCords(lat, lon float64) (model.Cities, error) {
	const op = "repositories.cities.GetByCoords"

	stmt := SELECT(Cities.ID, Cities.Name, Cities.Latitude, Cities.Longitude).FROM(Cities).WHERE(AND(Cities.Latitude.EQ(Float(lat)), Cities.Longitude.EQ(Float(lon)))).LIMIT(1)

	var city model.Cities
	err := stmt.Query(r.db, &city)
	if err != nil {
		return city, fmt.Errorf("%s: %w", op, err)
	}

	return city, nil
}

func (r *repository) FindByCords(lat, lon float64) (model.Cities, float64, error) {
	const op = "repositories.cities.FindByCords"

	var distance float64
	stmt, err := r.db.Prepare(`WITH distances AS (
        SELECT 
            name, 
            latitude, 
            longitude,
            6371 * ACOS(
                COS(RADIANS($1)) * COS(RADIANS(latitude)) *
                COS(RADIANS(longitude) - RADIANS($2)) +
                SIN(RADIANS($1)) * SIN(RADIANS(latitude))
            ) AS distance
        FROM cities
    )
    SELECT name, latitude, longitude, distance
    FROM distances
    ORDER BY distance
    LIMIT 1;`)
	if err != nil {
		return model.Cities{}, distance, fmt.Errorf("%s: %w", op, err)
	}

	var city model.Cities
	err = stmt.QueryRow(lat, lon).Scan(&city.Name, &city.Latitude, &city.Longitude, &distance)
	if err != nil {
		return city, distance, fmt.Errorf("%s: %w", op, err)
	}

	return city, distance, nil
}
