package services

import (
	"deprimera/api/daos"
	"deprimera/api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RouterAsistentes(e *echo.Echo) {
	e.POST("/api/asistentes", SaveAsistente)
	e.DELETE("/api/asistentes/:id_asistente/:id_persona", DeleteAsistente)
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
