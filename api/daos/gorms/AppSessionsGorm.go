package gorms

import (
	"database/sql"
	"time"
)

type AppSessionsGorm struct {
	ExpireDate time.Time      `gorm:"column:expire_date"`
	Token      sql.NullString `gorm:"column:token"`
	UserID     string         `gorm:"column:user_id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (a *AppSessionsGorm) TableName() string {
	return "app_sessions"
}
