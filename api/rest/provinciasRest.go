package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"

	"github.com/labstack/echo/v4"
)

func RouterProvincias(e *echo.Echo) {
	e.GET("/api/provincias", GetProvincias)
	e.GET("/api/provincias/:idPais/:idProvincia", GetPais)
}

func GetProvincias(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	provincias := daos.GetProvinciasDao().GetAll()
	return c.JSON(http.StatusOK, provincias)
}

func GetProvincia(c echo.Context) error {
	idPais, err := strconv.Atoi(c.Param("idPais"))
	if err != nil {
		log.Panic(err)
	}

	idProvincia, err := strconv.Atoi(c.Param("idProvincia"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	provinciaGorm := daos.GetProvinciasDao().Get(idPais, idProvincia)
	provincia := &models.Provincias{}
	copier.Copy(&provincia, &provinciaGorm)
	return c.JSON(http.StatusOK, provincia)
}
