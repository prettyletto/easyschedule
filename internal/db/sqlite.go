package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Init() *sqlx.DB {
	var DB *sqlx.DB
	var err error

	log.Println("Establishing connection with DataBase")
	DB, err = sqlx.Connect("sqlite3", "app.db")
	if err != nil {
		log.Fatalf("Failed to open a connection with the DataBase: %v", err)
	}

	DB.SetMaxOpenConns(1)

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to connect with the DataBase: %v", err)
	}
	log.Println("DataBase connected successfully")
	RunMigrations("app.db", "/Projects/easychedule/db/migrations/")

	return DB
}
