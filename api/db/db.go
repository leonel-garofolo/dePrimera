package db

import (
	"github.com/jinzhu/gorm"
)

type DB struct {
	Client *gorm.DB
}

func Get(connStr string) (*DB, error) {
	db, err := get(connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

func (d *DB) Close() error {
	return d.Client.Close()
}

func get(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
