package dbconnect

import (
	"go-server/config"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open(
		config.DatabaseDriver,
		config.DatabaseURL,
	)
	if err != nil {
		log.Println(err)
	}
	return db
}

func Get(query string, args ...any) (*sql.Rows, error) {
	db := Connect()
	defer db.Close()
	rows, err := db.Query(query, args)
	if err != nil {
		log.Println(err)
	}
	return rows, err
}

func Insert(query string, args ...any) error {
	db := Connect()
	defer db.Close()
	_, err := db.Exec(query, args)
	if err != nil {
		log.Println(err)
	}
	return err
}

func Update(query string, args ...any) error {
	db := Connect()
	defer db.Close()
	_, err := db.Exec(query, args)
	if err != nil {
		log.Println(err)
	}
	return err
}

func Delete(query string, args ...any) error {
	db := Connect()
	defer db.Close()
	_, err := db.Exec(query, args)
	if err != nil {
		log.Println(err)
	}
	return err
}
