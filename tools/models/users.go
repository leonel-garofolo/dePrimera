package models

type users struct {
	Clave       null.String `gorm:"column:clave" json:"clave"`
	Descripcion null.String `gorm:"column:descripcion" json:"descripcion"`
	Email       null.String `gorm:"column:email" json:"email"`
	Habilitado  `gorm:"column:habilitado" json:"habilitado"`
	Telefono    null.String `gorm:"column:telefono" json:"telefono"`
	UserID      string      `gorm:"column:user_id;primary_key" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (u *users) TableName() string {
	return "users"
}

// TableName sets the insert table name for this struct type
func (P *users) GetP() string {
	return "users"
}
