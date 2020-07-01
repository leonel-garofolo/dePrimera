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

func RouterPersonas(e *echo.Echo) {
	e.GET("/api/personas", GetPersonas)
	e.GET("/api/personas/:id", GetPersona)
	e.POST("/api/personas", SavePersona)
	e.DELETE("/api/personas/:id", DeletePersona)
	e.GET("/api/personas/info", InfoPersonas)
}

func GetPersonas(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	personas := daos.GetPersonasDao().GetAll()
	return c.JSON(http.StatusOK, personas)
}

func GetPersona(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	persona := daos.GetPersonasDao().Get(id)
	return c.JSON(http.StatusOK, persona)
}

func SavePersona(c echo.Context) error {
	personas := &models.Personas{}
	c.Bind(personas)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetPersonasDao().Save(personas)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
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
