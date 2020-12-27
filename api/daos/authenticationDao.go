package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
)

// ArbitrosDaoImpl struct
type AuthenticationDaoImpl struct{}

// Login user
func (ed *AuthenticationDaoImpl) Login(user string, pass string) gorms.UsersGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	userGorm := gorms.UsersGorm{}
	row := db.QueryRow("select * from app_users where id_user = ? and clave = ?", user, pass)
	error := row.Scan(&userGorm.IDUser, &userGorm.Clave, &userGorm.Nombre, &userGorm.Apellido, &userGorm.Habilitado, &userGorm.Telefono)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
		userGorm.IDUser = ""
	}
	return userGorm
}

// Register usuario
func (ed *AuthenticationDaoImpl) Register(user *gorms.UsersGorm) string {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("insert into app_users (id_user, clave, nombre, apellido, habilitado, telefono) values(?,?,?,?,?,?)",
		user.IDUser, user.Clave, user.Nombre, user.Apellido, user.Habilitado, user.Telefono)
	if error != nil {
		log.Println(error)
		panic(error)
	}
	return user.IDUser
}

// Reset password of user
func (ed *AuthenticationDaoImpl) ResetPassword(idUser string, oldPassword string, newPassword string) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("update app_users set clave =? where id_user =? and clave=?", newPassword, idUser, oldPassword)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return true
}
