package application

import (
	"deprimera/src/config"
	"deprimera/src/db"
	"fmt"

	"github.com/jinzhu/gorm"
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

func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"root",
		"de_primera_app",
	))
	if err != nil {
		return nil, err
	}
	return db, nil
}
