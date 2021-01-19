package daos

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
)

// CampeonatosDaoImpl struct
type CampeonatosDaoImpl struct{}

// GetAll campeonatos
func (ed *CampeonatosDaoImpl) GetAll() []gorms.CampeonatosGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select id_campeonato, id_liga, id_modelo, descripcion, fecha_inicio, fecha_fin from campeonatos")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var campeonatos []gorms.CampeonatosGorm
	for rows.Next() {
		campeonato := gorms.CampeonatosGorm{}
		error := rows.Scan(&campeonato.IDCampeonato, &campeonato.IDLiga, &campeonato.IDModelo, &campeonato.Descripcion, &campeonato.FechaInicio, &campeonato.FechaFin)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		campeonatos = append(campeonatos, campeonato)
	}
	return campeonatos
}

func (ed *CampeonatosDaoImpl) GetCampeonatoForUser(idUser string, idGrupo int) []gorms.CampeonatosGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	//IdGrupo is Admin
	query := ""
	switch idGrupo {
	case 1: //ADMINISTRADOR
		query = " select c.id_campeonato, c.id_liga, c.id_modelo, c.descripcion, c.fecha_inicio, c.fecha_fin " +
			" from campeonatos c "
		break
	case 2: //DELEGADOS
		query = "select c.id_campeonato, c.id_liga, c.id_modelo, c.descripcion, c.fecha_inicio, c.fecha_fin " +
			"from campeonatos c " +
			"inner join campeonatos_equipos ce on ce.id_campeonato = c.id_campeonato " +
			"inner join asistentes a on a.id_campeonato = ce.id_campeonato " +
			"inner join personas p on p.id_persona = a.id_persona " +
			"where p.id_user= '" + idUser + "' and p.idgrupo = " + strconv.Itoa(idGrupo)
		break
	case 3: //JUGADORES
		query = "select c.id_campeonato, c.id_liga, c.id_modelo, c.descripcion, c.fecha_inicio, c.fecha_fin " +
			"from campeonatos c " +
			"inner join campeonatos_equipos ce on ce.id_campeonato = c.id_campeonato " +
			"inner join jugadores j on j.id_equipo = ce.id_equipo " +
			"inner join personas p on p.id_persona = j.id_persona " +
			"where p.id_user= '" + idUser + "' and p.idgrupo = " + strconv.Itoa(idGrupo)
		break
	case 4: //ARBITROS
		query = "select c.id_campeonato, c.id_liga, c.id_modelo, c.descripcion, c.fecha_inicio, c.fecha_fin " +
			"from campeonatos c " +
			"inner join campeonatos_equipos ce on ce.id_campeonato = c.id_campeonato " +
			"inner join arbitros a on a.id_campeonato = ce.id_campeonato " +
			"inner join personas p on p.id_persona = a.id_persona " +
			"where p.id_user= '" + idUser + "' and p.idgrupo = " + strconv.Itoa(idGrupo)
		break
	}

	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var campeonatos []gorms.CampeonatosGorm
	for rows.Next() {
		campeonato := gorms.CampeonatosGorm{}
		error := rows.Scan(&campeonato.IDCampeonato, &campeonato.IDLiga, &campeonato.IDModelo, &campeonato.Descripcion, &campeonato.FechaInicio, &campeonato.FechaFin)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		campeonatos = append(campeonatos, campeonato)
	}
	return campeonatos
}

// Get campeonatos
func (ed *CampeonatosDaoImpl) Get(id int) gorms.CampeonatosGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select id_campeonato, id_liga, id_modelo, descripcion, fecha_inicio, fecha_fin from campeonatos where id_campeonato = ?", id)
	campeonato := gorms.CampeonatosGorm{}
	error := row.Scan(&campeonato.IDCampeonato, &campeonato.IDLiga, &campeonato.IDModelo, &campeonato.Descripcion, &campeonato.FechaInicio, &campeonato.FechaFin)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return campeonato
}

// Save campeonatos
func (ed *CampeonatosDaoImpl) Save(e *gorms.CampeonatosGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	fechaInicio := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		e.FechaInicio.Year(), e.FechaInicio.Month(), e.FechaInicio.Day(),
		e.FechaInicio.Hour(), e.FechaInicio.Minute(), e.FechaInicio.Second())
	fechaFin := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		e.FechaFin.Year(), e.FechaFin.Month(), e.FechaFin.Day(),
		e.FechaFin.Hour(), e.FechaFin.Minute(), e.FechaFin.Second())
	log.Println(fechaInicio)
	log.Println(fechaFin)
	if e.IDCampeonato > 0 {

		_, error := db.Exec("update campeonatos set descripcion=?, fecha_fin=?, fecha_inicio=?, id_liga=?, id_modelo=? where id_campeonato = ?",
			e.Descripcion, fechaFin, fechaInicio, e.IDLiga, e.IDModelo, e.IDCampeonato)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into campeonatos"+
			" (descripcion, fecha_fin, fecha_inicio, id_campeonato, id_liga, id_modelo) "+
			" values(?,?,?,?,?,?)", e.Descripcion, fechaFin, fechaInicio, e.IDCampeonato, e.IDLiga, e.IDModelo)
		IDCampeonato, error := res.LastInsertId()

		if error != nil {
			log.Println(error)
			panic(error)
		}
		e.IDCampeonato = IDCampeonato
	}
	return e.IDCampeonato
}

// Delete campeonatos
func (ed *CampeonatosDaoImpl) Delete(id int64) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from campeonatos where id_campeonato = ?", id)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			return false, error
		}
	}
	return true, nil
}
