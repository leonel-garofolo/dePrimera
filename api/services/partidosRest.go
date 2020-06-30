package services

import (
	"deprimera/api/daos"
	"deprimera/api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RouterPartidos(e *echo.Echo) {
	e.GET("/api/partidos", GetPartidos)
	e.GET("/api/partidos/:id", GetPartido)
	e.POST("/api/partidos", SavePartido)
	e.DELETE("/api/partidos", DeletePartido)
	e.GET("/api/partidos/info", InfoPartidos)
}

func GetPartidos(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	partidos := daos.GetPartidosDao().GetAll()
	return c.JSON(http.StatusOK, partidos)
}

func GetPartido(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	partido := daos.GetPartidosDao().Get(id)
	return c.JSON(http.StatusOK, partido)
}

func SavePartido(c echo.Context) error {
	partidos := &models.Partidos{}
	c.Bind(partidos)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetPartidosDao().Save(partidos)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeletePartido(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetPartidosDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
}

func InfoPartidos(c echo.Context) error {
	partidos := &models.Partidos{}
	c.Bind(partidos)

	j, err := json.Marshal(partidos)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, partidos)
	}
}
