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

func RouterEliminatorias(e *echo.Echo) {
	e.GET("/api/eliminatorias", GetEliminatorias)
	e.GET("/api/eliminatorias/:id", GetEliminatoria)
	e.POST("/api/eliminatorias", SaveEliminatoria)
	e.DELETE("/api/eliminatorias/:id", DeleteEliminatoria)
	e.GET("/api/eliminatorias/info", InfoEliminatorias)
}

func GetEliminatorias(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	eliminatoriasGorm := daos.GetEliminatoriasDao().GetAll()
	eliminatorias := []models.Eliminatorias{}
	copier.Copy(&eliminatorias, &eliminatoriasGorm)
	return c.JSON(http.StatusOK, eliminatorias)
}

func GetEliminatoria(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	eliminatoriaGorm := daos.GetEliminatoriasDao().Get(id)
	eliminatoria := &models.Eliminatorias{}
	copier.Copy(&eliminatoria, &eliminatoriaGorm)
	return c.JSON(http.StatusOK, eliminatoria)
}

func SaveEliminatoria(c echo.Context) error {
	eliminatorias := &models.Eliminatorias{}
	c.Bind(eliminatorias)

	eliminatoriasGorm := &gorms.EliminatoriasGorm{}
	copier.Copy(&eliminatoriasGorm, &eliminatorias)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetEliminatoriasDao().Save(eliminatoriasGorm)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteEliminatoria(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetEliminatoriasDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
}

func InfoEliminatorias(c echo.Context) error {
	eliminatorias := &models.Eliminatorias{}
	c.Bind(eliminatorias)

	j, err := json.Marshal(eliminatorias)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, eliminatorias)
	}
}
