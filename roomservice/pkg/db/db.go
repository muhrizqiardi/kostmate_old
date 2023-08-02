package db

import "github.com/jmoiron/sqlx"

type IDB interface {
	InsertNewRoomQuery(unique_name string)
}

type DB struct {
	db sqlx.DB
}

func NewDB(db sqlx.DB) *DB {
	return &DB{db: db}
}
