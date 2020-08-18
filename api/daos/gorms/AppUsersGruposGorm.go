package gorms

import "database/sql"

type AppUsersGruposGorm struct {
	IDGrupo sql.NullInt64 `gorm:"column:id_grupo"`
	UserID  string        `gorm:"column:user_id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (a *AppUsersGruposGorm) TableName() string {
	return "app_users_grupos"
}
