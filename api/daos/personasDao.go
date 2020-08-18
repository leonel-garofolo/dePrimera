package daos

import (
	"database/sql"
	"deprimera/api/application"
	"deprimera/api/daos/gorms"
	"log"
)

type PersonasDaoImpl struct{}

func (ed *PersonasDaoImpl) GetAll() []gorms.PersonasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select id_persona, apellido_nombre, domicilio, edad, id_liga, id_localidad, id_pais, id_provincia, id_tipo_doc, nro_doc from personas")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	personas := []gorms.PersonasGorm{}
	for rows.Next() {
		persona := gorms.PersonasGorm{}
		error := rows.Scan(&persona.IDPersona, &persona.ApellidoNombre, &persona.Domicilio, &persona.Edad, &persona.IDLiga, &persona.IDLocalidad, &persona.IDPais, &persona.IDProvincia, &persona.IDTipoDoc, &persona.NroDoc)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		personas = append(personas, persona)
	}
	return personas
}

func (ed *PersonasDaoImpl) Get(id int) gorms.PersonasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select id_persona, apellido_nombre, domicilio, edad, id_liga, id_localidad, id_pais, id_provincia, id_tipo_doc, nro_doc from personas where id_persona = ?", id)
	persona := gorms.PersonasGorm{}
	error := row.Scan(&persona.IDPersona, &persona.ApellidoNombre, &persona.Domicilio, &persona.Edad, &persona.IDLiga, &persona.IDLocalidad, &persona.IDPais, &persona.IDProvincia, &persona.IDTipoDoc, &persona.NroDoc)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return persona
}

func (ed *PersonasDaoImpl) Save(e *gorms.PersonasGorm) int64 {
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
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into personas"+
			" (id_persona, apellido_nombre, domicilio, edad, id_liga, id_localidad, id_pais, id_provincia, id_tipo_doc, nro_doc) "+
			" values(?,?,?,?,?,?,?,?,?,?)", e.IDPersona, e.ApellidoNombre, e.Domicilio, e.Edad, e.IDLiga, e.IDLocalidad, e.IDPais, e.IDProvincia, e.IDTipoDoc, e.NroDoc)

		IDPersona, error := res.LastInsertId()

		if error != nil {
			log.Println(error)
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
		log.Println(error)
		panic(error)
	}
	return true
}
