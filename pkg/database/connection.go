package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func Connect() (*sql.DB, error) {
	dsn := " host=fullstack-postgres port=5432 user=postgres dbname=fullstack_api sslmode=disable password=mysecretpassword";
	db, err := sql.Open("postgres", dsn);
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil;
}