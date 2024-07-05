package session

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"kiwi/.gen/kiwi/public/model"
	. "kiwi/.gen/kiwi/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

type Repository interface {
	Get(tg_id int64) (model.Session, error)
	Set(tg_id int64, value model.Session) error
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

func (r *repository) Get(tg_id int64) (model.Session, error) {
	const op = "repositories.session.Get"

	stmt := SELECT(Users.Session).FROM(Users).WHERE(Users.TelegramID.EQ(Int64(tg_id)))

	var user model.Users

	err := stmt.Query(r.db, &user)

	session := user.Session
	if err != nil {
		return session, fmt.Errorf("%s: %w", op, err)
	}

	return session, nil
}

func (r *repository) Set(tg_id int64, value model.Session) error {
	const op = "repositories.session.Set"

	stmt := Users.UPDATE(Users.Session).SET(value).WHERE(Users.TelegramID.EQ(Int64(tg_id)))

	_, err := stmt.Exec(r.db)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
