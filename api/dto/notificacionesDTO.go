package models

type Notificaciones struct {
	IDNotificacion int64  `json:"id_notificacion"`
	Titulo         string `json:"titulo"`
	Texto          string `json:"texto"`
	IDGrupo        int64  `json:"id_grupo"`
}

// TableName sets the insert table name for this struct type
func (a *Notificaciones) TableName() string {
	return "arbitros"
}
