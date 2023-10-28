package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB interface {
	Open() error
	Close() error
}

type DB struct {
	db *sql.DB
}

func (d *DB) Open() error {
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("DB connection established.")

	_, err = db.Exec(createSchema)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Create schema executed successfully.")

	d.db = db

	// Is this needed?
	defer db.Close()

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
