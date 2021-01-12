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

func RouterZonas(e *echo.Echo) {
	e.GET("/api/zonas", GetZonas)
	e.GET("/api/zonas/:id", GetZona)
	e.POST("/api/zonas", SaveZona)
	e.DELETE("/api/zonas/:id", DeleteZona)
	e.GET("/api/zonas/info", InfoZonas)
}

func GetZonas(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	zonasGorm := daos.GetZonasDao().GetAll()
	zonas := []models.Zonas{}
	copier.Copy(&zonas, &zonasGorm)
	return c.JSON(http.StatusOK, zonas)
}

func GetZona(c echo.Context) error {
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

func SaveZona(c echo.Context) error {
	zonas := &models.Zonas{}
	c.Bind(zonas)

	zonasGorm := &gorms.ZonasGorm{}
	copier.Copy(&zonasGorm, &zonas)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetZonasDao().Save(zonasGorm)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteZona(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	status, error := daos.GetZonasDao().Delete(id)

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

func InfoZonas(c echo.Context) error {
	zonas := &models.Zonas{}
	c.Bind(zonas)

	j, err := json.Marshal(zonas)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, zonas)
	}
}
