package database

import (
	"database/sql"
	"fmt"
	"go-server/config"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	TaskName    string `json:"task_name"`
	Description string `json:"description"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Connect() *sql.DB {
	db, err := sql.Open(config.DatabaseDriver, config.DatabaseURL)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func Get(query string) (*sql.Rows, error) {
	db := Connect()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}

func Insert(query string, args []string) error {
	db := Connect()
	defer db.Close()

	_, err := db.Exec(query, args)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func Update(query string, args []string) error {
	db := Connect()
	defer db.Close()

	_, err := db.Exec(query, args)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func Delete(query string, args []string) error {
	db := Connect()
	defer db.Close()

	_, err := db.Exec(query, args)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
