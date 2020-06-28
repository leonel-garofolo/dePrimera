package router

import (
	"deprimera/api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Mi Home")
}

func NewRouter(e *echo.Echo) {
	services.RouterEquipos(e)
	services.RouterLigas(e)
	e.GET("/api/home", home)
}
