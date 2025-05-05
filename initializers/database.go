package initializers

import (
	"database/sql"
	"fmt"
  "os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {

  var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s " + 
    "password=%s dbname=%s sslmode=disable", 
    os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), 
    os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), 
    os.Getenv("DB_NAME"))
  DB, err = sql.Open("postgres", psqlInfo)

  if err != nil {
    panic(err)
  }

  fmt.Println("Database Successfully connected!\n")

}

func DisconnectDB() {

  defer DB.Close()
  err := DB.Ping()

  if err != nil {
    panic(err)
  }

  fmt.Println("Database Successfully disconnected!\n")

}