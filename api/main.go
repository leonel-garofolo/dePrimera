package main

import (
	"net/http"
	"os"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("enviroment.yaml")
	port := os.Getenv("PORT")
	if os.Args != nil && len(os.Args) > 1 {
		env := os.Args[1]
		test := "test"
		if env == test {
			viper.SetConfigFile("enviroment-local.yaml")
			port = "8081"
		}

	}
	/*
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Error while reading config file %s", err)
		}
	*/

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) //permite cualquier dominio

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Bienvenido a la API de De Primera")
	})

	router.NewRouter(e)
	e.Logger.Fatal(e.Start(":" + port))
}
