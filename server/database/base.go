package database

import(
  "fmt"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
  db, err := sql.Open("sqlite3", "./server/database/data.db")
  if err != nil {
    fmt.Println(err)
  }
  return db
}
