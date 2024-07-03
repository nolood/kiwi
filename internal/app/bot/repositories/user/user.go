package user

import (
	"errors"
	"fmt"
	"kiwi/.gen/kiwi/public/model"
	. "kiwi/.gen/kiwi/public/table"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jmoiron/sqlx"
	"github.com/mymmrac/telego"
	"go.uber.org/zap"
)

type Repository interface {
	Get(tg_id int64) (model.Users, error)
	Create(user *telego.User) (model.Users, error)
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

func (r *repository) Get(tg_id int64) (model.Users, error) {
	const op = "repositories.user.Get"

	var user model.Users

	stmt := SELECT(Users.AllColumns).FROM(Users).WHERE(Users.TelegramID.EQ(Int64(tg_id)))

	err := stmt.Query(r.db, &user)
	if err != nil {

		if errors.Is(err, qrm.ErrNoRows) {
			return user, nil
		}

		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (r *repository) Create(tgUser *telego.User) (model.Users, error) {
	const op = "repositories.user.Create"

	newUser := model.Users{
		TelegramID:   tgUser.ID,
		IsPremium:    tgUser.IsPremium,
		LanguageCode: &tgUser.LanguageCode,
		FirstName:    &tgUser.FirstName,
		LastName:     &tgUser.LastName,
		Username:     tgUser.Username,
	}

	stmt := Users.INSERT(Users.FirstName, Users.LastName, Users.Username, Users.TelegramID, Users.IsPremium, Users.LanguageCode).MODEL(newUser).ON_CONFLICT(Users.TelegramID).DO_NOTHING().RETURNING(Users.AllColumns)

	var user model.Users

	err := stmt.Query(r.db, &user)
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
