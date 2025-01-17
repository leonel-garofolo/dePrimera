package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/dto/response"

	"github.com/labstack/echo/v4"
)

func RouterPartidos(e *echo.Echo) {
	e.GET("/api/partidos", GetPartidos)
	e.GET("/api/partidos/date/:date", GetPartidosFromDate)
	e.GET("/api/partidos/equipo/:id_equipo", GetPartidosFromEquipo)
	e.POST("/api/partidos/result", SaveResult)
	e.GET("/api/partidos/:id", GetPartido)
	e.POST("/api/partidos", SavePartido)
	e.DELETE("/api/partidos/:id", DeletePartido)
	e.GET("/api/partidos/info", InfoPartidos)
	e.GET("/api/partidos/history/:date", HistoryPartido)
	e.GET("/api/partidos/dates", GetFutureDatesPartidos)
}

func GetPartidos(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	partidosGorm := daos.GetPartidosDao().GetAll()
	partidos := []models.Partidos{}
	copier.Copy(&partidos, &partidosGorm)
	return c.JSON(http.StatusOK, partidos)
}

func GetPartidosFromDate(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	partidosFromDate := daos.GetPartidosDao().GetAllFromDate(c.Param("date"))
	return c.JSON(http.StatusOK, partidosFromDate)
}

func GetPartidosFromEquipo(c echo.Context) error {
	idEquipo, err := strconv.Atoi(c.Param("id_equipo"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	partidosFromDate := daos.GetPartidosDao().GetAllFromEquipo(idEquipo)
	return c.JSON(http.StatusOK, partidosFromDate)
}

func GetPartido(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	partidoGorm := daos.GetPartidosDao().Get(id)
	partido := &models.Partidos{}
	copier.Copy(&partido, &partidoGorm)
	return c.JSON(http.StatusOK, partido)
}

func SavePartido(c echo.Context) error {
	partidos := &models.Partidos{}
	c.Bind(partidos)

	partidosGorm := &gorms.PartidosGorm{}
	copier.Copy(&partidosGorm, &partidos)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetPartidosDao().Save(partidosGorm)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func SaveResult(c echo.Context) error {
	partidos := &models.PartidoResult{}
	c.Bind(partidos)

	partidosGorm := &gorms.PartidoResultGorm{}
	copier.Copy(&partidosGorm, &partidos)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetPartidosDao().SaveResult(partidosGorm)

	log.Println(id)
	if id > 0 {
		updated, error := daos.GetSancionesDao().SavePartido(partidosGorm.IDPartidos, partidosGorm.SancionAmarillasLocal, partidosGorm.SancionRojasLocal, partidosGorm.SancionAmarillasVisitante, partidosGorm.SancionRojasVisitante)
		if error == nil && updated {
			status, error := daos.GetSancionesDao().SavePartidoFinalizado(partidosGorm.IDPartidos, partidosGorm.Finalizado)
			fmt.Println("finalizado")
			fmt.Println(status)
			if error != nil {
				fmt.Println(error)
			}

			status, error = daos.GetCampeonatosDao().SaveCampeonatosGoleadores(partidosGorm.IDPartidos, partidos.GoleadoresLocal, partidos.GoleadoresVisitante)
			fmt.Println("goleadores")
			fmt.Println(status)
			if error != nil {
				fmt.Println(error)
			}
		} else {
			fmt.Println(error)
		}

	}
	return c.String(http.StatusOK, "insertado")
}

func DeletePartido(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	status, error := daos.GetPartidosDao().Delete(id)

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
	return c.String(http.StatusOK, "delete")
}

func FinalizarPartido(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	partidoGorm := daos.GetPartidosDao().Get(id)

	statusResult := "GL"
	if partidoGorm.ResultadoLocal < partidoGorm.ResultadoVisitante {
		statusResult = "GV"
	} else if partidoGorm.ResultadoLocal == partidoGorm.ResultadoVisitante {
		statusResult = "E"
	}

	daos.GetPartidosDao().FinalizarPartido(
		partidoGorm.IDLiga,
		partidoGorm.IDCampeonato,
		partidoGorm.IDEquipoLocal,
		partidoGorm.IDEquipoVisitante,
		statusResult)
	return c.String(http.StatusOK, "finalizado")
}

func InfoPartidos(c echo.Context) error {
	partidos := &models.Partidos{}
	c.Bind(partidos)

	j, err := json.Marshal(partidos)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, partidos)
	}
}

func HistoryPartido(c echo.Context) error {
	fromEquipo, err := strconv.Atoi(c.Param("date"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	partidosFromDateGorm := daos.GetPartidosDao().HistoryPlays(fromEquipo)
	partidosFromDate := []models.PartidosFromDate{}
	copier.Copy(&partidosFromDate, &partidosFromDateGorm)

	return c.JSON(http.StatusOK, partidosFromDate)
}

func GetFutureDatesPartidos(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	datesPartidos := daos.GetPartidosDao().GetFuturePartidos()
	return c.JSON(http.StatusOK, datesPartidos)
}
