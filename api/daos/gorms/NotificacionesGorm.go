package gorms

import "database/sql"

type NotificacionesGorm struct {
	IDGrupo        sql.NullInt64 `gorm:"column:id_grupo"`
	IDNotificacion int64           `gorm:"column:id_notificacion;primary_key"`
	Texto          string        `gorm:"column:texto"`
	Titulo         string        `gorm:"column:titulo"`
}

// TableName sets the insert table name for this struct type
func (n *NotificacionesGorm) TableName() string {
	return "notificaciones"
}
