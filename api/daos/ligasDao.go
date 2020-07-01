package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type LigasDaoImpl struct{}

func (ed *LigasDaoImpl) GetAll() []models.Ligas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from ligas")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var ligas []models.Ligas
	for rows.Next() {
		liga := models.Ligas{}
		rows.Scan(&liga.Cuit)
		rows.Scan(&liga.Domicilio)
		rows.Scan(&liga.IDLiga)
		rows.Scan(&liga.MailContacto)
		rows.Scan(&liga.Nombre)
		rows.Scan(&liga.NombreContacto)
		rows.Scan(&liga.Telefono)
		rows.Scan(&liga.TelefonoContacto)
		ligas = append(ligas, liga)
	}
	return ligas
}

func (ed *LigasDaoImpl) Get(id int) models.Ligas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	liga := models.Ligas{}
	row := db.QueryRow("select * from ligas where id_liga = ?", id)
	row.Scan(&liga.Cuit)
	row.Scan(&liga.Domicilio)
	row.Scan(&liga.IDLiga)
	row.Scan(&liga.MailContacto)
	row.Scan(&liga.Nombre)
	row.Scan(&liga.NombreContacto)
	row.Scan(&liga.Telefono)
	row.Scan(&liga.TelefonoContacto)

	return liga
}

func (ed *LigasDaoImpl) Save(e *models.Ligas) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDLiga > 0 {
		_, error := db.Exec("update ligas"+
			" set cuit=?, domicilio=?, id_liga=?, mail_contacto=?, nombre=?, nombre_contacto=?, telefono=?, telefono_contacto=? "+
			" where id_liga = ?", e.Cuit, e.Domicilio, e.MailContacto, e.Nombre, e.NombreContacto, e.Telefono, e.TelefonoContacto, e.IDLiga)

		if error != nil {
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into ligas"+
			" (cuit, domicilio, id_liga, mail_contacto, nombre, nombre_contacto, telefono, telefono_contacto) "+
			" values(?,?,?,?,?,?,?,?)", e.Cuit, e.Domicilio, e.IDLiga, e.MailContacto, e.Nombre, e.NombreContacto, e.Telefono, e.TelefonoContacto)

		idEquipo, error := res.LastInsertId()

		if error != nil {
			panic(error)
		}
		e.IDLiga = idEquipo
	}
	return e.IDLiga
}

func (ed *LigasDaoImpl) Delete(id int) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from ligas where id_liga = ?", id)
	if error != nil {
		panic(error)
	}
	return true
}
