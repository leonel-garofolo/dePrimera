package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/dto/response"
	help "github.com/leonel-garofolo/dePrimeraApiRest/api/help"

	"github.com/labstack/echo/v4"
)

func RouterCampeonatos(e *echo.Echo) {
	e.GET("/api/campeonatos", GetCampeonatos)
	e.GET("/api/campeonatos/:id", GetCampeonato)
	e.POST("/api/campeonatos", SaveCampeonato)
	e.DELETE("/api/campeonatos/:id", DeleteCampeonato)
	e.GET("/api/campeonatos/fixture/:id_campeonato", GetFixture)
	e.GET("/api/campeonatos/goleadores/:id_campeonato", GetGoleadores)
	e.GET("/api/campeonatos/user/:id_user/:id_grupo", GetCampeonatosForUserID)
	e.GET("/api/campeonatos/table/:id_campeonato", GetTablePosition)
	e.GET("/api/campeonatos/sanciones/:id_campeonato", GetJugadoresSanciones)
	e.POST("/api/campeonatos/fixture/generate/:id_liga/:id_campeonato", GenerateFixture)
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
	if id > 0 {
		if campeonatos.GenFixture && !campeonatos.GenFixtureFinish {
			error := generateFixture(c, campeonatos.IDLiga, int(campeonatos.IDCampeonato))
			fmt.Println(error)
		}
	}
	return c.String(http.StatusOK, "insertado")
}

func DeleteCampeonato(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	status, error := daos.GetCampeonatosDao().Delete(id)

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

	log.Println(id)
	return c.JSON(http.StatusOK, resp)
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
	partidosFromDate := daos.GetPartidosDao().GetAllFromCampeonato(idTorneo)
	return c.JSON(http.StatusOK, partidosFromDate)
}

func GetCampeonatosForUserID(c echo.Context) error {
	idGrupo, err := strconv.Atoi(c.Param("id_grupo"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	campeonatos := daos.GetCampeonatosDao().GetCampeonatoForUser(c.Param("id_user"), idGrupo)
	return c.JSON(http.StatusOK, campeonatos)
}

func GetTablePosition(c echo.Context) error {
	idTorneo, err := strconv.Atoi(c.Param("id_campeonato"))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("id_campeonato: ")
	fmt.Println(idTorneo)

	daos := daos.NewDePrimeraDaos()
	equiposTablePos := daos.GetPartidosDao().GetTablePosition(idTorneo)
	return c.JSON(http.StatusOK, equiposTablePos)
}

func GetJugadoresSanciones(c echo.Context) error {
	idTorneo, err := strconv.Atoi(c.Param("id_campeonato"))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("id_campeonato: ")
	fmt.Println(idTorneo)

	daos := daos.NewDePrimeraDaos()
	sancionesJugadoresFromCampeonato := daos.GetSancionesDao().GetSancionesFromCampeonato(idTorneo)
	return c.JSON(http.StatusOK, sancionesJugadoresFromCampeonato)
}

func GenerateFixture(c echo.Context) error {
	idLiga, errLiga := strconv.Atoi(c.Param("id_liga"))
	if errLiga != nil {
		log.Panic(errLiga)
	}

	idTorneo, err := strconv.Atoi(c.Param("id_campeonato"))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("id_liga: " + c.Param("id_liga") + " | id_campeonato: " + c.Param("id_campeonato"))

	return generateFixture(c, idLiga, idTorneo)
}

func generateFixture(c echo.Context, idLiga int, idTorneo int) error {
	daos := daos.NewDePrimeraDaos()
	equiposGorm := daos.GetEquiposDao().GetAllFromCampeonato(idLiga, idTorneo)

	fixtureService := help.FixtureHelp{}
	fixture := fixtureService.CalcularLiga(len(equiposGorm))

	daos.GetEquiposDao().UpdateNro(idLiga, idTorneo)

	daos.GetPartidosDao().SaveFixture(idLiga, idTorneo, time.Now(), fixture)
	daos.GetPartidosDao().FinishFixtureGen(idLiga, idTorneo)

	return c.JSON(http.StatusOK, fixture)
}

func GetGoleadores(c echo.Context) error {
	idTorneo, err := strconv.Atoi(c.Param("id_campeonato"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	goleadores := daos.GetCampeonatosDao().GetGoleadores(idTorneo)

	return c.JSON(http.StatusOK, goleadores)
}
