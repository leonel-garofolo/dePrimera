package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"

	"github.com/labstack/echo/v4"
)

func RouterNotificaciones(e *echo.Echo) {
	e.GET("/api/notificaciones", GetNotificaciones)
	e.POST("/api/notificaciones", SaveNotificacion)
	e.DELETE("/api/notificaciones/:id_notificacion", DeleteNotificacion)
	e.GET("/api/notificaciones/info", InfoNotificaciones)
}

func GetNotificaciones(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	notificacionesGorm := daos.GetNotificacionesDao().GetAll()

	notificaciones := []models.Notificaciones{}
	copier.Copy(&notificaciones, &notificacionesGorm)

	return c.JSON(http.StatusOK, notificaciones)
}

func SaveNotificacion(c echo.Context) error {
	notificaciones := &models.Notificaciones{}
	c.Bind(notificaciones)

	notificacionesGorm := &gorms.NotificacionesGorm{}
	copier.Copy(&notificacionesGorm, &notificaciones)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetNotificacionesDao().Save(notificacionesGorm)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteNotificacion(c echo.Context) error {
	idNotificacion, err := strconv.ParseInt(c.Param("id_notificacion"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	daos.GetNotificacionesDao().Delete(idNotificacion)

	log.Println(idNotificacion)
	return c.String(http.StatusOK, "delete")
}

func InfoNotificaciones(c echo.Context) error {
	notificaciones := &models.Notificaciones{}
	c.Bind(notificaciones)

	j, err := json.Marshal(notificaciones)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, notificaciones)
	}
}
