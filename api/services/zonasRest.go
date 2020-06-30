package services

import (
	"deprimera/api/daos"
	"deprimera/api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RouterZonas(e *echo.Echo) {
	e.GET("/api/zonas", GetZonas)
	e.GET("/api/zonas/:id", GetZona)
	e.POST("/api/zonas", SaveZona)
	e.DELETE("/api/zonas", DeleteZona)
	e.GET("/api/zonas/info", InfoZonas)
}

func GetZonas(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	zonas := daos.GetZonasDao().GetAll()
	return c.JSON(http.StatusOK, zonas)
}

func GetZona(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	zona := daos.GetZonasDao().Get(id)
	return c.JSON(http.StatusOK, zona)
}

func SaveZona(c echo.Context) error {
	zonas := &models.Zonas{}
	c.Bind(zonas)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetZonasDao().Save(zonas)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteZona(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetZonasDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
}

func InfoZonas(c echo.Context) error {
	zonas := &models.Zonas{}
	c.Bind(zonas)

	j, err := json.Marshal(zonas)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, zonas)
	}
}
