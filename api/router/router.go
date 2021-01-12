package router

import (
	"net/http"

	services "github.com/leonel-garofolo/dePrimeraApiRest/api/rest"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Mi Home")
}

func NewRouter(e *echo.Echo) {
	services.RouterArbitros(e)
	services.RouterAsistentes(e)
	services.RouterJugadores(e)
	services.RouterCampeonatos(e)
	services.RouterEliminatorias(e)
	services.RouterEquipos(e)
	services.RouterLigas(e)
	services.RouterPartidos(e)
	services.RouterPersonas(e)
	services.RouterSanciones(e)
	services.RouterZonas(e)
	services.RouterNotificaciones(e)
	services.RouterAppGrupos(e)
	services.RouterAuthentication(e)

	services.RouterPaises(e)
	services.RouterProvincias(e)
	services.RouterQuery(e)
	services.RouterComentarios(e)
	e.GET("/api/home", home)
}
