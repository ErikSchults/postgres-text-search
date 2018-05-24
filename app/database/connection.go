package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func New() *sql.DB {
	addr := "postgres://textsearch:supersecret@db_textsearch/db_textsearch?sslmode=disable"
	conn, err := sql.Open("postgres", addr)

	if err != nil {
		panic(err)
	}

	waitForConnection(conn)

	return conn
}

func GetConnection() *sql.DB {
	if conn == nil {
		return New()
	}
	return conn
}

func waitForConnection(db *sql.DB) {
	err := db.Ping()

	for i := 0; i < 5; i++ {
		if err == nil {
			fmt.Println("connected to database")
			return
		}
		log.Println(err)
		time.Sleep(time.Duration(i*2) * time.Second)
	}

	panic(err)
}
