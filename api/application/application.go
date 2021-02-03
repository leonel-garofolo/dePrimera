package application

import (
	"database/sql"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/config"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/db"

	_ "github.com/lib/pq"
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
	/*
			printConn := fmt.Sprintf(
				"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
				viper.Get("database.user"),
				viper.Get("database.pass"),
				viper.Get("database.type"),
				viper.Get("database.host"),
				viper.Get("database.port"),
				viper.Get("database.name"),
			)
			fmt.Println(printConn)

		printConn := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			viper.Get("database.user"),
			viper.Get("database.pass"),
			viper.Get("database.host"),
			viper.Get("database.port"),
			viper.Get("database.name"),
		)
		fmt.Println(printConn)
	*/
	db, err := sql.Open("postgres", "postgres://kfaneklolueulx:838051822695796f4ba0a1764a70c02cca4a71abeda6a08af177389f260966c3@ec2-54-221-220-82.compute-1.amazonaws.com:5432/d6afam55njg5u6")
	if err != nil {
		return nil, err
	}
	return db, nil
}
