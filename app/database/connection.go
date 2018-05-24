package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	DBCon *sql.DB
)

func New() *sql.DB {
	addr := "postgres://textsearch:supersecret@db_textsearch/db_textsearch?sslmode=disable"
	db, err := sql.Open("postgres", addr)

	if err != nil {
		panic(err)
	}

	waitForConnection(db)

	return db
}

func waitForConnection(db *sql.DB) {
	err := db.Ping()

	for i := 0; i < 5; i++ {
		if err == nil {
			log.Println("Database connection established")
			return
		}
		log.Println(err)
		time.Sleep(time.Duration(i*2) * time.Second)
	}

	panic(err)
}
