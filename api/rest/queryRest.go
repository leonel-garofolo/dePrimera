package services

import (
	"net/http"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"

	"github.com/labstack/echo/v4"
)

func RouterQuery(e *echo.Echo) {
	e.GET("/api/query/configuraciones", GetQueryConfiguraciones)
}

func GetQueryConfiguraciones(c echo.Context) error {

	daos := daos.NewDePrimeraDaos()
	query := daos.GetQueryDao().GetConfiguracionesSize()
	return c.JSON(http.StatusOK, query)
}
