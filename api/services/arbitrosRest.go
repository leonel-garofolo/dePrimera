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

func RouterArbitros(e *echo.Echo) {
	e.GET("/api/arbitros", GetArbitros)
	e.GET("/api/arbitros/:id", GetArbitro)
	e.POST("/api/arbitros", SaveArbitro)
	e.DELETE("/api/arbitros", DeleteArbitro)
	e.GET("/api/arbitros/info", InfoArbitros)
}

func GetArbitros(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	arbitros := daos.GetArbitrosDao().GetAll()
	return c.JSON(http.StatusOK, arbitros)
}

func GetArbitro(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	arbitro := daos.GetArbitrosDao().Get(id)
	return c.JSON(http.StatusOK, arbitro)
}

func SaveArbitro(c echo.Context) error {
	arbitros := &models.Arbitros{}
	c.Bind(arbitros)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetArbitrosDao().Save(arbitros)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func DeleteArbitro(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetArbitrosDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
}

func InfoArbitros(c echo.Context) error {
	arbitros := &models.Arbitros{}
	c.Bind(arbitros)

	j, err := json.Marshal(arbitros)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, arbitros)
	}
}
