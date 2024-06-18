package main

import (
	"fmt"
	"kiwi/internal/config"
	"kiwi/internal/storage/postgres"

	"kiwi/.gen/kiwi/public/model"
	. "kiwi/.gen/kiwi/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

func main() {

	cfg := config.MustLoad()

	// log := logger.New(cfg.Env)

	storage := postgres.New(cfg.Storage)

	var users []model.Users

	stmt := SELECT(Users.ID).FROM(Users)

	stmt.Query(storage.Db, &users)

	fmt.Println(users)

}
