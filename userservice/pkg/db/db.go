package db

import (
	"github.com/jmoiron/sqlx"
	// "../common"
)

type IDB interface {
	InsertNewUserQuery()
}

type DB struct {
	db *sqlx.DB
}

func NewDB(db *sqlx.DB) *DB {
	return &DB{db}
}

func (d *DB) InsertNewUserQuery(email)
