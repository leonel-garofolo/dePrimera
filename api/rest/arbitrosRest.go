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

func RouterArbitros(e *echo.Echo) {
	e.GET("/api/arbitros", GetArbitros)
	e.POST("/api/arbitros", SaveArbitro)
	e.DELETE("/api/arbitros/:id_arbitro/:id_persona/id_campeonato", DeleteArbitro)
	e.GET("/api/arbitros/info", InfoArbitros)
}

func GetArbitros(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	arbitros := daos.GetArbitrosDao().GetAll()
	return c.JSON(http.StatusOK, arbitros)
}

func SaveArbitro(c echo.Context) error {
	arbitros := &models.Arbitros{}
	c.Bind(arbitros)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetArbitrosDao().Save(arbitros)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteArbitro(c echo.Context) error {
	idArbitro, err1 := strconv.ParseInt(c.Param("id_arbitro"), 10, 64)
	if err1 != nil {
		log.Panic(err1)
	}

	idPersona, err2 := strconv.ParseInt(c.Param("id_persona"), 10, 64)
	if err2 != nil {
		log.Panic(err2)
	}

	idCampeonato, err2 := strconv.ParseInt(c.Param("id_campeonato"), 10, 64)
	if err2 != nil {
		log.Panic(err2)
	}
	daos := daos.NewDePrimeraDaos()
	status, error := daos.GetArbitrosDao().Delete(idArbitro, idPersona, idCampeonato)

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

	log.Println(idArbitro, idPersona, idCampeonato)
	return c.JSON(http.StatusOK, resp)
}

func InfoArbitros(c echo.Context) error {
	arbitros := &models.Arbitros{}
	c.Bind(arbitros)

	j, err := json.Marshal(arbitros)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, arbitros)
	}
}
