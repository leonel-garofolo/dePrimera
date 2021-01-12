package services

import (
	"log"
	"net/http"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"

	"github.com/labstack/echo/v4"
)

func RouterComentarios(e *echo.Echo) {
	e.POST("/api/comentarios", SaveComentario)
}

func SaveComentario(c echo.Context) error {
	comentarios := &models.Comentarios{}
	c.Bind(comentarios)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetComentariosDao().Save(comentarios)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}
