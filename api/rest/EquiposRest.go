package services

import (
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
	"github.com/jinzhu/copier"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"github.com/labstack/echo/v4"
)

func RouterEquipos(e *echo.Echo) {
	e.GET("/api/equipos", GetEquipos)
	e.GET("/api/equipos/:id", GetEquipo)
	e.POST("/api/equipos", SaveEquipos)
	e.DELETE("/api/equipos/:id", DeleteEquipo)
	e.GET("/api/equipos/info", InfoEquipo)
}

func GetEquipos(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	equiposGorm := daos.GetEquiposDao().GetAll()
	equipos := []models.Equipos{}
	copier.Copy(&equipos, &equiposGorm)
	return c.JSON(http.StatusOK, equipos)
}

func GetEquipo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	equipo := daos.GetEquiposDao().Get(id)
	return c.JSON(http.StatusOK, equipo)
}

func SaveEquipos(c echo.Context) error {
	equipos := &models.Equipos{}
	c.Bind(equipos)

	equiposGorm := &gorms.EquiposGorm{}
	copier.Copy(&equiposGorm, &equipos)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetEquiposDao().Save(equiposGorm)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteEquipo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetEquiposDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
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
