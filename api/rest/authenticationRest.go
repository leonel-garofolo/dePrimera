package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
	models "github.com/leonel-garofolo/dePrimeraApiRest/api/dto"
)

func RouterAuthentication(e *echo.Echo) {
	e.POST("/api/authentication/login", Login)
	e.POST("/api/authentication/register", Register)
	e.POST("/api/authentication/forgot", Forgot)
	e.POST("/api/authentication/reset", ResetPassword)
	e.GET("/api/authentication/permiso/:user", GetUserAppGrupos)

}

func Login(c echo.Context) error {
	user := &models.Users{}
	c.Bind(user)

	daos := daos.NewDePrimeraDaos()
	userGorm := daos.GetAuthenticationDao().Login(user.UserID, user.Password)
	user.UserID = userGorm.UserID
	user.Telefono = userGorm.Telefono.String
	user.Habilitado = userGorm.Habilitado

	return c.JSON(http.StatusOK, user)
}

func Register(c echo.Context) error {
	user := &models.Users{}
	c.Bind(user)

	userGorm := &gorms.UsersGorm{}
	copier.Copy(&userGorm, &user)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetAuthenticationDao().Register(userGorm)

	log.Println(id)
	return c.String(http.StatusOK, "insertado")
}

func Forgot(c echo.Context) error {
	email := c.Param("id")
	fmt.Println(email)
	// TODO send mail

	return c.JSON(http.StatusOK, "mail enviado")
}

func ResetPassword(c echo.Context) error {
	user := &models.UsersPassReset{}
	c.Bind(user)

	daos := daos.NewDePrimeraDaos()
	id := daos.GetAuthenticationDao().ResetPassword(user.Email, user.OldPassword, user.NewPassword)
	fmt.Println(id)
	return c.JSON(http.StatusOK, "reset password")
}

func GetUserAppGrupos(c echo.Context) error {
	userId := c.Param("user")
	daos := daos.NewDePrimeraDaos()
	gruposGorms := daos.GetAppGruposDao().GetUserAppGrupos(userId)

	grupos := models.AppGrupos{}
	copier.Copy(&grupos, &gruposGorms)

	return c.JSON(http.StatusOK, grupos)
}
