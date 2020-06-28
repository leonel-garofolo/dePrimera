package services

import (
	"database/sql"
	"deprimera/api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RouterLigas(e *echo.Echo) {
	e.GET("/api/ligas", GetLigas)
	e.POST("/api/ligas", SaveLiga)
	e.GET("/api/ligas/info", InfoLigas)
}

func GetLigas(c echo.Context) error {
	ligas := models.Ligas{}
	return c.JSON(http.StatusOK, ligas)
}

func SaveLiga(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}

	ligas := &models.Ligas{
		Cuit: sql.NullString{
			String: m["cuit"].(string),
			Valid:  false,
		},
		Nombre: m["nombre"].(string),
	}

	ligas.SaveLigas()
	log.Println("ligas id : " + string(ligas.IDLiga))
	return c.String(http.StatusOK, "insertado")
}

func InfoLigas(c echo.Context) error {
	ligas := &models.Ligas{}
	c.Bind(ligas)

	j, err := json.Marshal(ligas)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, ligas)
	}
}
