package daos

import (
	"database/sql"
	"deprimera/api/application"
	"deprimera/api/daos/gorms"
	"log"
)

type LigasDaoImpl struct{}

func (ed *LigasDaoImpl) GetAll() []gorms.LigasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from ligas")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	ligas := []gorms.LigasGorm{}
	for rows.Next() {
		liga := gorms.LigasGorm{}
		error := rows.Scan(&liga.IDLiga, &liga.Nombre, &liga.NombreContacto, &liga.MailContacto, &liga.Cuit, &liga.Domicilio, &liga.Telefono, &liga.TelefonoContacto)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		ligas = append(ligas, liga)
	}
	return ligas
}

func (ed *LigasDaoImpl) Get(id int) gorms.LigasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	liga := gorms.LigasGorm{}
	row := db.QueryRow("select * from ligas where id_liga = ?", id)
	error := row.Scan(&liga.IDLiga, &liga.Cuit, &liga.Domicilio, &liga.MailContacto, &liga.Nombre, &liga.NombreContacto, &liga.Telefono, &liga.TelefonoContacto)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return liga
}

func (ed *LigasDaoImpl) Save(e *gorms.LigasGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDLiga > 0 {
		_, error := db.Exec("update ligas"+
			" set cuit=?, domicilio=?, mail_contacto=?, nombre=?, nombre_contacto=?, telefono=?, telefono_contacto=? "+
			" where id_liga = ?", e.Cuit, e.Domicilio, e.MailContacto, e.Nombre, e.NombreContacto, e.Telefono, e.TelefonoContacto, e.IDLiga)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into ligas (cuit, domicilio, mail_contacto, nombre, nombre_contacto, telefono, telefono_contacto) "+
			" values(?,?,?,?,?,?,?)", e.Cuit, e.Domicilio, e.MailContacto, e.Nombre, e.NombreContacto, e.Telefono, e.TelefonoContacto)
		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDLiga, _ = res.LastInsertId()
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
