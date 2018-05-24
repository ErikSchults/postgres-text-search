package database

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func Migrate() {
	m := createMigrator()

	runMigrations(m)
}

func createMigrator() *migrate.Migrate {
	driver, err := postgres.WithInstance(DBCon, &postgres.Config{})
	if err != nil {
		panic("Failed to initialize postgres instance: " + err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+os.Getenv("GO_PROJECT_DIR")+"/migrations",
		"db_textsearch", driver,
	)
	if err != nil {
		panic("Failed to initialize postgres instance: " + err.Error())
	}
	return m
}

func runMigrations(m *migrate.Migrate) {
	err := m.Up()

	switch err {
	case nil:
		log.Println("Database migrated successfully")
	case migrate.ErrNoChange:
		log.Println("Database already up-to-date")
	default:
		panic("Failed to run migrations: " + err.Error())
	}
}
