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

func RouterPaises(e *echo.Echo) {
	e.GET("/api/paises", GetPaises)
	e.GET("/api/paises/:id", GetPais)
}

func GetPaises(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	paises := daos.GetPaisesDao().GetAll()
	return c.JSON(http.StatusOK, paises)
}

func GetPais(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	zonaGorm := daos.GetPaisesDao().Get(id)
	zona := &models.Paises{}
	copier.Copy(&zona, &zonaGorm)
	return c.JSON(http.StatusOK, zona)
}
