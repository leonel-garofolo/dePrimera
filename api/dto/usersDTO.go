package models

type Users struct {
	UserID     string `json:"user_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Telefono   string `json:"telefono"`
	Habilitado bool   `json:"habilitado"`
}

type UsersPassReset struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// TableName sets the insert table name for this struct type
func (u *Users) TableName() string {
	return "users"
}
