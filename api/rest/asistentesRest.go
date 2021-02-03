package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/dto/response"

	"github.com/labstack/echo/v4"
)

func RouterAsistentes(e *echo.Echo) {
	e.GET("/api/asistentes", GetAsistentes)
	e.POST("/api/asistentes", SaveAsistente)
	e.DELETE("/api/asistentes/:id_asistente/:id_persona/:id_campeonato", DeleteAsistente)
	e.GET("/api/asistentes/info", InfoAsistentes)
}

func GetAsistentes(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	asistentes := daos.GetAsistentesDao().GetAll()
	return c.JSON(http.StatusOK, asistentes)
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

	idCampeonato, err := strconv.ParseInt(c.Param("id_campeonato"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	status, error := daos.GetAsistentesDao().Delete(idAsistente, idPersona, idCampeonato)

	resp := &response.UpdatedResponse{}
	resp.Status = status
	if !status {
		resp.Message = "Error al intentar eliminar el Registro."
		sError := error.Error()
		fmt.Println(sError)
		if strings.Contains(sError, "Cannot") {
			resp.Message = "El registro no se pudo eliminar."
		}
	}

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
