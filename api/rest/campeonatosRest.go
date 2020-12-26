package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
	help "github.com/leonel-garofolo/dePrimeraApiRest/api/help"

	"github.com/labstack/echo/v4"
)

func RouterCampeonatos(e *echo.Echo) {
	e.GET("/api/campeonatos", GetCampeonatos)
	e.GET("/api/campeonatos/:id", GetCampeonato)
	e.POST("/api/campeonatos", SaveCampeonato)
	e.DELETE("/api/campeonatos/:id", DeleteCampeonato)
	e.GET("/api/campeonato/fixture/:id_campeonato", GetFixture)
	e.GET("/api/campeonato/fixture/generate", GetFixtureGenerate)
	e.GET("/api/campeonatos/info", InfoCampeonatos)
}

func GetCampeonatos(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	campeonatosGorm := daos.GetCampeonatosDao().GetAll()
	campeonatos := []models.Campeonatos{}
	copier.Copy(&campeonatos, &campeonatosGorm)
	return c.JSON(http.StatusOK, campeonatos)
}

func GetCampeonato(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	campeonatoGorm := daos.GetCampeonatosDao().Get(id)
	campeonato := &models.Campeonatos{}
	copier.Copy(&campeonato, &campeonatoGorm)
	return c.JSON(http.StatusOK, campeonato)
}

func SaveCampeonato(c echo.Context) error {
	campeonatos := &models.Campeonatos{}
	c.Bind(campeonatos)

	campeonatosGorm := &gorms.CampeonatosGorm{}
	copier.Copy(&campeonatosGorm, &campeonatos)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetCampeonatosDao().Save(campeonatosGorm)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteCampeonato(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetCampeonatosDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
}

func InfoCampeonatos(c echo.Context) error {
	campeonatos := &models.Campeonatos{}
	c.Bind(campeonatos)

	j, err := json.Marshal(campeonatos)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, campeonatos)
	}
}

func GetFixture(c echo.Context) error {
	idTorneo, err := strconv.Atoi(c.Param("id_campeonato"))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("id_campeonato: ")
	fmt.Println(idTorneo)

	daos := daos.NewDePrimeraDaos()
	campeonatoGorm := daos.GetCampeonatosDao().Get(idTorneo)
	equiposGorm := daos.GetEquiposDao().GetAllFromCampeonato(campeonatoGorm.IDLiga)

	fixtureService := help.FixtureHelp{}
	fixture := fixtureService.CalcularLiga(len(equiposGorm))

	daos.GetEquiposDao().UpdateNro(idTorneo)

	daos.GetPartidosDao().SaveFixture(1, 1, time.Now(), fixture)

	return c.JSON(http.StatusOK, fixture)
}

func GetFixtureGenerate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	zonaGorm := daos.GetZonasDao().Get(id)
	zona := &models.Zonas{}
	copier.Copy(&zona, &zonaGorm)
	return c.JSON(http.StatusOK, zona)
}
