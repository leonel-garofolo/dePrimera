package main

import (
	"deprimera/api/router"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("../app.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	miMain := viper.Get("test")
	log.Println(miMain)

	value := viper.Get("database.url")
	log.Println(value)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) //permite cualquier dominio

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World2!")
	})

	router.NewRouter(e)
	e.Logger.Fatal(e.Start(":8081"))
}
