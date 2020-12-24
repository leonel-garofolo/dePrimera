package services

import (
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
	"github.com/jinzhu/copier"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RouterLigas(e *echo.Echo) {
	e.GET("/api/ligas", GetLigas)
	e.GET("/api/ligas/:id", GetLiga)
	e.POST("/api/ligas", SaveLiga)
	e.DELETE("/api/ligas/:id", DeleteLiga)
	e.GET("/api/ligas/info", InfoLigas)
}

func GetLigas(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	ligasGorm := daos.GetLigasDao().GetAll()
	
	ligas := []models.Ligas{}
	copier.Copy(&ligas, &ligasGorm)

	return c.JSON(http.StatusOK, &ligas)
}

func GetLiga(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	
	ligaGorm := daos.GetLigasDao().Get(id)
	liga := &models.Ligas{}
	copier.Copy(&liga, &ligaGorm)

	return c.JSON(http.StatusOK, liga)
}

func SaveLiga(c echo.Context) error {
	ligas := &models.Ligas{}
	c.Bind(ligas)

	ligasGorm := &gorms.LigasGorm{}
	copier.Copy(&ligasGorm, &ligas)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetLigasDao().Save(ligasGorm)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteLiga(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetLigasDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
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