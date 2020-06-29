package services

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RouterEquipos(e *echo.Echo) {
	e.GET("/api/equipos", GetEquipos)
	e.POST("/api/equipos", SaveEquipos)
	e.GET("/api/equipos/info", InfoEquipo)
}

func GetEquipos(c echo.Context) error {
	equipos := models.GetAllEquipo()
	return c.JSON(http.StatusOK, equipos)
}

func SaveEquipos(c echo.Context) error {
	equipos := &models.Equipos{}
	c.Bind(equipos)

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())

	}
	db.Create(equipos)

	log.Println(equipos)
	return c.String(http.StatusOK, "insertado")
}

func InfoEquipo(c echo.Context) error {
	equipos := &models.Equipos{}
	c.Bind(equipos)

	j, err := json.Marshal(equipos)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, equipos)
	}
}
