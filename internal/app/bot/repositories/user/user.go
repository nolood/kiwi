package user

import (
	"errors"
	"fmt"
	"kiwi/.gen/kiwi/public/model"
	. "kiwi/.gen/kiwi/public/table"
	userdto "kiwi/internal/app/bot/dto/user"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jmoiron/sqlx"
	"github.com/mymmrac/telego"
	"go.uber.org/zap"
)

type Repository interface {
	Get(tgId int64) (userdto.UserWithProfile, error)
	Create(user *telego.User) (userdto.UserWithProfile, error)
	UpdateName(tgId int64, name string) error
	UpdateAge(tgId int64, age int) error
	UpdateGender(tgId int64, gender string) error
	UpdateAbout(tgId int64, about string) error
	UpdatePhoto(tgId int64, about string) error
	UpdateLocation(tgId int64, latitude *float64, longitude *float64) error
	UpdateCity(tgId int64, city string) error
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

func (r *repository) UpdateName(tgId int64, newValue string) error {
	const op = "bot.repositories.user.UpdateName"

	stmt := Profiles.UPDATE(Profiles.Name).SET(newValue).WHERE(Profiles.UserTgID.EQ(Int64(tgId)))
	_, err := stmt.Exec(r.db)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *repository) UpdateLocation(tgId int64, latitude *float64, longitude *float64) error {

	const op = "bot.repositories.user.UpdateLocation"

	stmt := Profiles.UPDATE(Profiles.Latitude, Profiles.Longitude).SET(latitude, longitude).WHERE(Profiles.UserTgID.EQ(Int64(tgId)))
	_, err := stmt.Exec(r.db)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *repository) Get(tgId int64) (userdto.UserWithProfile, error) {
	const op = "bot.repositories.user.Get"

	var userprof userdto.UserWithProfile

	var user model.Users
	var profile model.Profiles

	stmtUser := SELECT(Users.AllColumns).FROM(Users).WHERE(Users.TelegramID.EQ(Int64(tgId)))

	err := stmtUser.Query(r.db, &user)
	if err != nil {

		if errors.Is(err, qrm.ErrNoRows) {
			return userprof, nil
		}

		return userprof, fmt.Errorf("%s: %w", op, err)
	}

	userprof.User = user

	stmtProfile := SELECT(Profiles.AllColumns).FROM(Profiles).WHERE(Profiles.UserID.EQ(Int64(int64(user.ID))))

	err = stmtProfile.Query(r.db, &profile)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return userprof, nil
		}

		return userprof, fmt.Errorf("%s: %w", op, err)
	}

	userprof.Profile = profile

	return userprof, nil
}

func (r *repository) Create(tgUser *telego.User) (userdto.UserWithProfile, error) {
	const op = "bot.repositories.user.Create"

	var userprof userdto.UserWithProfile

	newUser := model.Users{
		TelegramID:   tgUser.ID,
		IsPremium:    tgUser.IsPremium,
		LanguageCode: &tgUser.LanguageCode,
		FirstName:    &tgUser.FirstName,
		LastName:     &tgUser.LastName,
		Username:     tgUser.Username,
	}

	stmtUser := Users.INSERT(Users.FirstName, Users.LastName, Users.Username, Users.TelegramID, Users.IsPremium, Users.LanguageCode).MODEL(newUser).ON_CONFLICT(Users.TelegramID).DO_NOTHING().RETURNING(Users.AllColumns)

	var user model.Users
	var profile model.Profiles

	err := stmtUser.Query(r.db, &user)
	if err != nil {

		if errors.Is(err, qrm.ErrNoRows) {
			return userprof, nil
		}

		return userprof, fmt.Errorf("%s: %w", op, err)
	}

	userprof.User = user

	stmtProfile := Profiles.INSERT(Profiles.UserID, Profiles.UserTgID).VALUES(user.ID, user.TelegramID).ON_CONFLICT(Profiles.UserID).DO_NOTHING().RETURNING(Profiles.AllColumns)

	err = stmtProfile.Query(r.db, &profile)
	if err != nil {

		if errors.Is(err, qrm.ErrNoRows) {
			return userprof, nil
		}

		return userprof, fmt.Errorf("%s: %w", op, err)
	}

	userprof.Profile = profile

	return userprof, nil
}

func (r *repository) UpdateAbout(tgId int64, newValue string) error {
	const op = "bot.repositories.user.UpdateAbout"

	stmt := Profiles.UPDATE(Profiles.About).SET(newValue).WHERE(Profiles.UserTgID.EQ(Int64(tgId)))
	_, err := stmt.Exec(r.db)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *repository) UpdateGender(tgId int64, newValue string) error {
	const op = "bot.repositories.user.UpdateGender"

	stmt := Profiles.UPDATE(Profiles.Gender).SET(newValue).WHERE(Profiles.UserTgID.EQ(Int64(tgId)))
	_, err := stmt.Exec(r.db)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *repository) UpdateAge(tgId int64, newValue int) error {
	const op = "bot.repositories.user.UpdateAge"

	stmt := Profiles.UPDATE(Profiles.Age).SET(newValue).WHERE(Profiles.UserTgID.EQ(Int64(tgId)))
	_, err := stmt.Exec(r.db)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *repository) UpdatePhoto(tgId int64, newValue string) error {
	const op = "bot.repositories.user.UpdatePhoto"

	stmt := Profiles.UPDATE(Profiles.PhotoID).SET(newValue).WHERE(Profiles.UserTgID.EQ(Int64(tgId)))
	_, err := stmt.Exec(r.db)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *repository) UpdateCity(tgId int64, newValue string) error {
	const op = "bot.repositories.user.UpdateCity"

	stmt := Profiles.UPDATE(Profiles.Location).SET(newValue).WHERE(Profiles.UserTgID.EQ(Int64(tgId)))

	_, err := stmt.Exec(r.db)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
