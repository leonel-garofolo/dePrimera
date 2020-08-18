package models

import (
	"database/sql"
	"time"
)

type app_sessions struct {
	ExpireDate time.Time      `gorm:"column:expire_date"`
	Token      sql.NullString `gorm:"column:token"`
	UserID     string         `gorm:"column:user_id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (a *app_sessions) TableName() string {
	return "app_sessions"
}
