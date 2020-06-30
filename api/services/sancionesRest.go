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

func RouterSanciones(e *echo.Echo) {
	e.GET("/api/sancions", GetSanciones)
	e.GET("/api/sancions/:id", GetSancion)
	e.POST("/api/sancions", SaveSancion)
	e.DELETE("/api/sancions", DeleteSancion)
	e.GET("/api/sancions/info", InfoSanciones)
}

func GetSanciones(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	sancions := daos.GetSancionesDao().GetAll()
	return c.JSON(http.StatusOK, sancions)
}

func GetSancion(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	sancion := daos.GetSancionesDao().Get(id)
	return c.JSON(http.StatusOK, sancion)
}

func SaveSancion(c echo.Context) error {
	sancion := &models.Sanciones{}
	c.Bind(sancion)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetSancionesDao().Save(sancion)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteSancion(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetSancionesDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
}

func InfoSanciones(c echo.Context) error {
	sancions := &models.Sanciones{}
	c.Bind(sancions)

	j, err := json.Marshal(sancions)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, sancions)
	}
}
