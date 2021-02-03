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
	row := db.QueryRow("select * from app_users where id_user = $1 and clave = $2;", user, pass)
	error := row.Scan(&userGorm.UserID, &userGorm.Password, &userGorm.Habilitado, &userGorm.Telefono)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
		userGorm.UserID = ""
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

	_, error := db.Exec("insert into app_users (id_user, clave, habilitado, telefono) values($1,$2,$3,$4,$5,$6)",
		user.UserID, user.Password, user.Habilitado, user.Telefono)
	if error != nil {
		log.Println(error)
		panic(error)
	}
	return user.UserID
}

// Reset password of user
func (ed *AuthenticationDaoImpl) ResetPassword(idUser string, oldPassword string, newPassword string) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("update app_users set clave =$1 where id_user =$2 and clave=$3", newPassword, idUser, oldPassword)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return true
}
