package models

type app_users_grupos struct {
	IDGrupo null.Int `gorm:"column:id_grupo" json:"id_grupo"`
	UserID  string   `gorm:"column:user_id;primary_key" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (a *app_users_grupos) TableName() string {
	return "app_users_grupos"
}

// TableName sets the insert table name for this struct type
func (P *app_users_grupos) GetP() string {
	return "app_users_grupos"
}
