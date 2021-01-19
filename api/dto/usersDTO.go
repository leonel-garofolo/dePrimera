package models

type Users struct {
	UserID     string `json:"id_user"`
	Password   string `json:"password"`
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
