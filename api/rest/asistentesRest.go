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
	e.POST("/api/asistentes", SaveAsistente)
	e.DELETE("/api/asistentes/:id_asistente/:id_persona", DeleteAsistente)
	e.GET("/api/asistentes/info", InfoAsistentes)
}

func GetAsistentes(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	arbitros := daos.GetAsistentesDao().GetAll()
	return c.JSON(http.StatusOK, arbitros)
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
	idAsistente, err := strconv.ParseInt(c.Param("id_asistente"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	idPersona, err := strconv.ParseInt(c.Param("id_persona"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	daos.GetAsistentesDao().Delete(idAsistente, idPersona)

	log.Println(idAsistente, idPersona)
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
