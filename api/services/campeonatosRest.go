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

func RouterCampeonatos(e *echo.Echo) {
	e.GET("/api/campeonatos", GetCampeonatos)
	e.GET("/api/campeonatos/:id", GetCampeonato)
	e.POST("/api/campeonatos", SaveCampeonato)
	e.DELETE("/api/campeonatos/:id", DeleteCampeonato)
	e.GET("/api/campeonatos/info", InfoCampeonatos)
}

func GetCampeonatos(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	campeonatos := daos.GetCampeonatosDao().GetAll()
	return c.JSON(http.StatusOK, campeonatos)
}

func GetCampeonato(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	campeonato := daos.GetCampeonatosDao().Get(id)
	return c.JSON(http.StatusOK, campeonato)
}

func SaveCampeonato(c echo.Context) error {
	campeonatos := &models.Campeonatos{}
	c.Bind(campeonatos)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetCampeonatosDao().Save(campeonatos)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteCampeonato(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetCampeonatosDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
}

func InfoCampeonatos(c echo.Context) error {
	campeonatos := &models.Campeonatos{}
	c.Bind(campeonatos)

	j, err := json.Marshal(campeonatos)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, campeonatos)
	}
}
