package db

import "database/sql"

type Store interface {
	Querier
}
type SQLStore struct {
	db *sql.DB
	*Queries
}

func SetupDB(db *sql.DB) (store Store) {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
