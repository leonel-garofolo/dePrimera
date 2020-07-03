package application

import (
	"database/sql"
	"deprimera/api/config"
	"deprimera/api/db"
	"fmt"

	"github.com/spf13/viper"
)

// Application holds commonly used app wide data, for ease of DI
type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

// Get captures env vars, establishes DB connection and keeps/returns
// reference to both
func Get() (*Application, error) {
	cfg := config.Get()

	db, err := db.Get(cfg.GetDBConnStr())
	if err != nil {
		return nil, err
	}

	return &Application{
		DB:  db,
		Cfg: cfg,
	}, nil
}

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.Get("database.user"),
		viper.Get("database.pass"),
		viper.Get("database.type"),
		viper.Get("database.host"),
		viper.Get("database.port"),
		viper.Get("database.name"),
	))
	if err != nil {
		return nil, err
	}
	return db, nil
}
