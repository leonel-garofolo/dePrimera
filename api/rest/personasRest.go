package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"

	"github.com/labstack/echo/v4"
)

func RouterPersonas(e *echo.Echo) {
	e.GET("/api/personas", GetPersonas)
	e.GET("/api/personas/:id", GetPersona)
	e.POST("/api/personas", SavePersona)
	e.DELETE("/api/personas/:id", DeletePersona)
	e.GET("/api/personas/info", InfoPersonas)
}

func GetPersonas(c echo.Context) error {
	daos := daos.NewDePrimeraDaos()
	personasGorm := daos.GetPersonasDao().GetAll()

	personas := []models.Personas{}
	for n := range personasGorm {
		persona := parseGson(personasGorm[n])
		personas = append(personas, persona)
	}

	return c.JSON(http.StatusOK, personas)
}

func GetPersona(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}

	daos := daos.NewDePrimeraDaos()
	personaGorm := daos.GetPersonasDao().Get(id)
	persona := parseGson(personaGorm)
	return c.JSON(http.StatusOK, persona)
}

func SavePersona(c echo.Context) error {
	persona := &models.Personas{}
	c.Bind(persona)

	personasGorm := parseJson(persona)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetPersonasDao().Save(personasGorm)

	log.Println(id)
	return c.String(http.StatusOK, strconv.FormatInt(id, 10))
}

func DeletePersona(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panic(err)
	}
	daos := daos.NewDePrimeraDaos()
	daos.GetPersonasDao().Delete(id)

	log.Println(id)
	return c.String(http.StatusOK, "delete")
}

func InfoPersonas(c echo.Context) error {
	personas := &models.Personas{}
	c.Bind(personas)

	j, err := json.Marshal(personas)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error al obtener la info")
	} else {
		log.Println(string(j))

		return c.JSON(http.StatusOK, personas)
	}
}

func parseJson(json *models.Personas) *gorms.PersonasGorm {
	return &gorms.PersonasGorm{
		IDPersona:      json.IDPersona,
		ApellidoNombre: json.ApellidoNombre,
		Domicilio: sql.NullString{
			String: json.Domicilio,
			Valid:  false,
		},
		Edad: sql.NullInt64{
			Int64: json.Edad,
			Valid: false,
		},
		Localidad:   json.Localidad,
		IDPais:      json.IDPais,
		IDProvincia: json.IDProvincia,
		IDTipoDoc:   json.IDTipoDoc,
		NroDoc:      json.NroDoc,
	}
}

func parseGson(dto gorms.PersonasGorm) models.Personas {
	return models.Personas{
		IDPersona:      dto.IDPersona,
		ApellidoNombre: dto.ApellidoNombre,
		Domicilio:      dto.Domicilio.String,
		Edad:           dto.Edad.Int64,
		Localidad:      dto.Localidad,
		IDPais:         dto.IDPais,
		IDProvincia:    dto.IDProvincia,
		IDTipoDoc:      dto.IDTipoDoc,
		NroDoc:         dto.NroDoc,
	}
}
