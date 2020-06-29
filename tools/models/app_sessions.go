package models

type app_sessions struct {
	ExpireDate null.Time   `gorm:"column:expire_date" json:"expire_date"`
	Token      null.String `gorm:"column:token" json:"token"`
	UserID     string      `gorm:"column:user_id;primary_key" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (a *app_sessions) TableName() string {
	return "app_sessions"
}

// TableName sets the insert table name for this struct type
func (P *app_sessions) GetP() string {
	return "app_sessions"
}
