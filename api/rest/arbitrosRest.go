package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"

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
	arbitrosGorm := daos.GetArbitrosDao().GetAll()

	arbitros := []models.Arbitros{}
	copier.Copy(&arbitros, &arbitrosGorm)

	return c.JSON(http.StatusOK, arbitros)
}

func SaveArbitro(c echo.Context) error {
	arbitros := &models.Arbitros{}
	c.Bind(arbitros)

	arbitrosGorm := &gorms.ArbitrosGorm{}
	copier.Copy(&arbitrosGorm, &arbitros)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetArbitrosDao().Save(arbitrosGorm)

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
	daos.GetArbitrosDao().Delete(idArbitro, idPersona, idCampeonato)

	log.Println(idArbitro, idPersona, idCampeonato)
	return c.String(http.StatusOK, "delete")
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
