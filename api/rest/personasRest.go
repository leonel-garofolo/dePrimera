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

func RouterPersonas(e *echo.Echo) {
	e.GET("/api/personas", GetPersonas)
	e.GET("/api/personas/:id", GetPersona)
	e.POST("/api/personas", SavePersona)
	e.DELETE("/api/personas/:id", DeletePersona)
	e.GET("/api/personas/info", InfoPersonas)
}

func GetPersonas(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	personasGorm := daos.GetPersonasDao().GetAll()
	personas := []models.Personas{}
	copier.Copy(&personas, &personasGorm)
	return c.JSON(http.StatusOK, personas)
}

func GetPersona(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	personaGorm := daos.GetPersonasDao().Get(id)
	persona := &models.Personas{}
	copier.Copy(&persona, &personaGorm)
	return c.JSON(http.StatusOK, persona)
}

func SavePersona(c echo.Context) error {
	personas := &models.Personas{}
	c.Bind(personas)

	personasGorm := &gorms.PersonasGorm{}
	copier.Copy(&personasGorm, &personas)


	daos := daos.NewDePrimeraDaos()
	id := daos.GetPersonasDao().Save(personasGorm)

	log.Println(id)
	return c.String(http.StatusOK, strconv.FormatInt(id, 10))
}

func DeletePersona(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetPersonasDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
}

func InfoPersonas(c echo.Context) error {
	personas := &models.Personas{}
	c.Bind(personas)

	j, err := json.Marshal(personas)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, personas)
	}
}
