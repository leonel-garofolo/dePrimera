package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/dto/response"
)

func RouterEquipos(e *echo.Echo) {
	e.GET("/api/equipos", GetEquipos)
	e.GET("/api/equipos/:id", GetEquipo)
	e.GET("/api/equipos/user/:id_user/:id_grupo", GetEquiposFromUser)
	e.GET("/api/equipos/plantel/:id_equipo", GetPlantel)
	e.POST("/api/equipos", SaveEquipos)
	e.DELETE("/api/equipos/:id", DeleteEquipo)
	e.GET("/api/equipos/info", InfoEquipo)
}

func GetEquipos(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	equipos := daos.GetEquiposDao().GetAll()
	return c.JSON(http.StatusOK, equipos)
}

func GetEquipo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	equipo := daos.GetEquiposDao().Get(id)
	return c.JSON(http.StatusOK, equipo)
}

func GetEquiposFromUser(c echo.Context) error {
	idGrupo, err := strconv.Atoi(c.Param("id_grupo"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	equipos := daos.GetEquiposDao().GetEquiposFromUser(c.Param("id_user"), idGrupo)
	return c.JSON(http.StatusOK, equipos)
}

func GetPlantel(c echo.Context) error {
	idEquipo, err := strconv.Atoi(c.Param("id_equipo"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	equipos := daos.GetEquiposDao().GetPlantel(idEquipo)
	return c.JSON(http.StatusOK, equipos)
}

func SaveEquipos(c echo.Context) error {
	equipos := &models.Equipos{}
	c.Bind(equipos)

	equiposGorm := &gorms.EquiposGorm{}
	copier.Copy(&equiposGorm, &equipos)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetEquiposDao().Save(equiposGorm)
	if id > 0 {
		daos.GetEquiposDao().SaveEquiposCampeonatos(equiposGorm.IDCampeonato, id)
	}

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteEquipo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	status, error := daos.GetEquiposDao().Delete(id)

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

func InfoEquipo(c echo.Context) error {
	equipos := &models.Equipos{}
	c.Bind(equipos)

	j, err := json.Marshal(equipos)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, equipos)
	}
}
