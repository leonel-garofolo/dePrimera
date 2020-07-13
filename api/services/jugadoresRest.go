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

func RouterJugadores(e *echo.Echo) {
	e.GET("/api/jugadores", GetJugadores)
	e.POST("/api/jugadores", SaveJugador)
	e.DELETE("/api/jugadores/:id_jugador/:id_persona", DeleteJugador)
	e.GET("/api/jugadores/info", InfoJugadores)
}

func GetJugadores(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	arbitros := daos.GetJugadoresDao().GetAll()
	return c.JSON(http.StatusOK, arbitros)
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
	daos.GetJugadoresDao().Delete(idJugador, idPersona)

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
