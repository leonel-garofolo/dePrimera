package router

import (
	services "deprimera/api/rest"
	"net/http"

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

	e.GET("/api/home", home)
}
