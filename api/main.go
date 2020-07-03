package main

import (
	"deprimera/api/router"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile("../app.yaml")
	if len(os.Args) > 0 {
		env := os.Args[1]
		test := "test"
		if env == test {
			viper.SetConfigFile("../app-test.yaml")
		}

	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

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
