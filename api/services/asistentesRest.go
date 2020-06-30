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

func RouterAsistentes(e *echo.Echo) {
	e.GET("/api/asistentes", GetAsistentes)
	e.GET("/api/asistentes/:id", GetAsistente)
	e.POST("/api/asistentes", SaveAsistente)
	e.DELETE("/api/asistentes", DeleteAsistente)
	e.GET("/api/asistentes/info", InfoAsistentes)
}

func GetAsistentes(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	asistentes := daos.GetAsistentesDao().GetAll()
	return c.JSON(http.StatusOK, asistentes)
}

func GetAsistente(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	asistente := daos.GetAsistentesDao().Get(id)
	return c.JSON(http.StatusOK, asistente)
}

func SaveAsistente(c echo.Context) error {
	asistentes := &models.Asistentes{}
	c.Bind(asistentes)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetAsistentesDao().Save(asistentes)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteAsistente(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetAsistentesDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
}

func InfoAsistentes(c echo.Context) error {
	asistentes := &models.Asistentes{}
	c.Bind(asistentes)

	j, err := json.Marshal(asistentes)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, asistentes)
	}
}
