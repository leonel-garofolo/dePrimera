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

func RouterJugadores(e *echo.Echo) {
	e.GET("/api/jugadores", GetJugadores)
	e.POST("/api/jugadores", SaveJugador)
	e.DELETE("/api/jugadores/:id_jugador/:id_persona", DeleteJugador)
	e.GET("/api/jugadores/info", InfoJugadores)
}

func GetJugadores(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	jugadores := daos.GetJugadoresDao().GetAll()
	return c.JSON(http.StatusOK, jugadores)
}

func SaveJugador(c echo.Context) error {
	jugadores := &models.Jugadores{}
	c.Bind(jugadores)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetJugadoresDao().Save(jugadores)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteJugador(c echo.Context) error {
	idJugador, err := strconv.ParseInt(c.Param("id_jugador"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	idPersona, err := strconv.ParseInt(c.Param("id_persona"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	status, error := daos.GetJugadoresDao().Delete(idJugador, idPersona)

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

	log.Println(idJugador, idPersona)
	return c.String(http.StatusOK, "delete")
}

func InfoJugadores(c echo.Context) error {
	jugadores := &models.Jugadores{}
	c.Bind(jugadores)

	j, err := json.Marshal(jugadores)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, jugadores)
	}
}
