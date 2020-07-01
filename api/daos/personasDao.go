package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"log"
)

type PersonasDaoImpl struct{}

func (ed *PersonasDaoImpl) GetAll() []models.Personas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select * from personas")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	var personas []models.Personas
	for rows.Next() {
		persona := models.Personas{}
		rows.Scan(&persona.IDPersona)
		rows.Scan(&persona.ApellidoNombre)
		rows.Scan(&persona.Domicilio)
		rows.Scan(&persona.Edad)
		rows.Scan(&persona.IDLiga)
		rows.Scan(&persona.IDLocalidad)
		rows.Scan(&persona.IDPais)
		rows.Scan(&persona.IDProvincia)
		rows.Scan(&persona.IDTipoDoc)
		rows.Scan(&persona.NroDoc)
		personas = append(personas, persona)
	}
	return personas
}

func (ed *PersonasDaoImpl) Get(id int) models.Personas {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select * from personas where id_persona = ?", id)
	persona := models.Personas{}
	row.Scan(&persona.IDPersona)
	row.Scan(&persona.ApellidoNombre)
	row.Scan(&persona.Domicilio)
	row.Scan(&persona.Edad)
	row.Scan(&persona.IDLiga)
	row.Scan(&persona.IDLocalidad)
	row.Scan(&persona.IDPais)
	row.Scan(&persona.IDProvincia)
	row.Scan(&persona.IDTipoDoc)
	row.Scan(&persona.NroDoc)
	return persona
}

func (ed *PersonasDaoImpl) Save(e *models.Personas) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDPersona > 0 {
		_, error := db.Exec("update personas"+
			" set apellido_nombre=?, domicilio=?, edad=?, id_liga=?, id_localidad=?, id_pais=?, id_provincia=?, id_tipo_doc=?, nro_doc=? "+
			" where id_persona = ?", e.ApellidoNombre, e.Domicilio, e.Edad, e.IDLiga, e.IDLocalidad, e.IDPais, e.IDProvincia, e.IDTipoDoc, e.NroDoc, e.IDPersona)

		if error != nil {
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into personas"+
			" (id_persona, apellido_nombre, domicilio, edad, id_liga, id_localidad, id_pais, id_provincia, id_tipo_doc, nro_doc) "+
			" values(?,?,?,?,?,?,?,?,?,?)", e.IDPersona, e.ApellidoNombre, e.Domicilio, e.Edad, e.IDLiga, e.IDLocalidad, e.IDPais, e.IDProvincia, e.IDTipoDoc, e.NroDoc)

		IDPersona, error := res.LastInsertId()

		if error != nil {
			panic(error)
		}
		e.IDPersona = IDPersona
	}
	return e.IDPersona
}

func (ed *PersonasDaoImpl) Delete(id int) bool {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from personas where id_persona = ?", id)
	if error != nil {
		panic(error)
	}
	return true
}
