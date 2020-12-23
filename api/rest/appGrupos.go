package services

import (
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
	"github.com/jinzhu/copier"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RouterAppGrupos(e *echo.Echo) {
	e.GET("/api/app_grupos", GetAppGrupos)
	e.GET("/api/app_grupos/info", InfoAppGrupos)
}

func GetAppGrupos(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	gruposGorms := daos.GetAppGruposDao().GetAll()

	grupos := []models.AppGrupos{}
	copier.Copy(&grupos, &gruposGorms)

	return c.JSON(http.StatusOK, grupos)
}

func InfoAppGrupos(c echo.Context) error {
	grupos := &models.AppGrupos{}
	c.Bind(grupos)

	j, err := json.Marshal(grupos)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, grupos)
	}
}
