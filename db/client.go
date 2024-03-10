package db

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Pg *sqlx.DB
}

func Init(Pg *sqlx.DB) *DB {
	return &DB{
		Pg: Pg,
	}
}
